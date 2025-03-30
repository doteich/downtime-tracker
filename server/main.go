package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbPool                                                           *pgxpool.Pool
	Logger                                                           *slog.Logger
	pgHost, pgUser, pgPassword, pgDatabase, basicUser, basicPassword string
	AdminCreds                                                       map[string]string
)

type Event struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Start    time.Time `json:"startDate"`
	End      time.Time `json:"endDate"`
	Type     string    `json:"type"`
	Location string    `json:"location"`
	Color    string    `json:"color"`
}

type Location struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Id    int    `json:"id"` // This is the primary key
}

type DownTimeType struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Id    int    `json:"id"` // This is the primary key
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {

	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	flag.StringVar(&pgHost, "host", "localhost", "Postgres host")
	flag.StringVar(&pgUser, "user", "postgres", "Postgres user")
	flag.StringVar(&pgPassword, "password", "pass", "Postgres password")
	flag.StringVar(&pgDatabase, "db", "downtimes", "Postgres database")

	// pgHost = os.Getenv("PG_HOST")
	// pgUser = os.Getenv("PG_USER")
	// pgPassword = os.Getenv("PG_PASSWORD")
	// pgDatabase = os.Getenv("PG_DATABASE")
	basicUser = os.Getenv("BASIC_USER")
	basicPassword = os.Getenv("BASIC_PASSWORD")

	flag.Parse()

}

func main() {
	// Initialize the database connection pool
	if err := initDB(); err != nil {
		Logger.Error("Error initializing database connection pool: " + err.Error())
		return
	}
	defer dbPool.Close()
	creds := make(map[string]string)
	creds[basicUser] = basicPassword

	AdminCreds = make(map[string]string)

	if err := InitUsers(); err != nil {
		Logger.Error("Error initializing users: " + err.Error())
		return
	}

	// Create a new Chi router
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Group(func(r chi.Router) {

		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if !VerifyHeader(r) {
					http.Error(w, "Invalid username or password", http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
			})
		})
		r.Post("/events", CreateEvent)
		r.Delete("/events/{id}", DeleteEvent)
		r.Post("/locations", CreateLocation)
		r.Post("/downtime_types", CreateDowntimeType)

	})

	r.Group(func(r chi.Router) {
		//r.Use(middleware.BasicAuth("downtime", creds))
		r.Get("/events", GetEvents)
		r.Get("/locations", GetLocations)
		r.Get("/downtime_types", GetDowntimeTypes)
		r.Post("/login/hash", CreatePasswordHash)
		r.Post("/login", Login)
	})
	r.Group(func(r chi.Router) {
		r.Get("/healthz", CheckHealth)

	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe("127.0.0.1:"+port, r); err != nil {
		slog.Error("Error starting server: " + err.Error())
	}
}

func initDB() error {
	var err error

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", pgUser, pgPassword, pgHost, pgDatabase)
	dbPool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return err
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func InitUsers() error {
	rows, err := dbPool.Query(context.Background(), "SELECT username, password FROM public.users")

	if err != nil {
		return err
	}
	for rows.Next() {
		var username, password string
		if err := rows.Scan(&username, &password); err != nil {
			return err
		}
		AdminCreds[username] = password

	}

	return nil
}

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		Logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO events (name, start_date, end_date, location, color,type) VALUES ($1, $2, $3, $4, $5, $6)", event.Name, event.Start, event.End, event.Location, event.Color, event.Type)
	if err != nil {
		Logger.Error("Error inserting event: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	if startDate == "" || endDate == "" {
		http.Error(w, "start_date and end_date query parameters are required", http.StatusBadRequest)
		return
	}

	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		http.Error(w, "start_date must be in RFC3339 format", http.StatusBadRequest)
		return
	}

	end, err := time.Parse(time.RFC3339, endDate)

	if err != nil {
		http.Error(w, "end_date must be in RFC3339 format", http.StatusBadRequest)
		return
	}

	rows, err := dbPool.Query(context.Background(), "SELECT name, start_date, end_date, location, color, type, id FROM events WHERE start_date >= $1 AND end_date < $2", start, end)
	if err != nil {
		Logger.Error("Error querying events: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	events := make([]Event, 0)
	for rows.Next() {
		var event Event

		if err := rows.Scan(&event.Name, &event.Start, &event.End, &event.Location, &event.Color, &event.Type, &event.Id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		Logger.Error("Error scanning rows: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := dbPool.Exec(context.Background(), "DELETE FROM events WHERE id = $1", id)
	if err != nil {
		Logger.Error("Error deleting event: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		Logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO locations (name, color) VALUES ($1, $2)", location.Name, location.Color)
	if err != nil {
		Logger.Error("Error inserting location: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	rows, err := dbPool.Query(context.Background(), "SELECT id, name, color FROM locations")
	if err != nil {
		Logger.Error("Error querying locations: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	locations := make([]Location, 0)
	for rows.Next() {
		var location Location
		if err := rows.Scan(&location.Id, &location.Name, &location.Color); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		Logger.Error("Error scanning rows: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(locations)
}
func CreateDowntimeType(w http.ResponseWriter, r *http.Request) {
	var downtimeType DownTimeType
	if err := json.NewDecoder(r.Body).Decode(&downtimeType); err != nil {
		Logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO downtime_types (name, color) VALUES ($1, $2)", downtimeType.Name, downtimeType.Color)
	if err != nil {
		Logger.Error("Error inserting downtime type: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetDowntimeTypes(w http.ResponseWriter, r *http.Request) {
	rows, err := dbPool.Query(context.Background(), "SELECT id, name, color FROM downtime_types")
	if err != nil {
		Logger.Error("Error querying downtime types: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	downtimeTypes := make([]DownTimeType, 0)
	for rows.Next() {
		var downtimeType DownTimeType
		if err := rows.Scan(&downtimeType.Id, &downtimeType.Name, &downtimeType.Color); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		downtimeTypes = append(downtimeTypes, downtimeType)
	}

	if err := rows.Err(); err != nil {
		Logger.Error("Error scanning rows: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(downtimeTypes)
}

func VerifyHeader(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	pw, ok := AdminCreds[username]
	if !ok {
		return false
	}

	err := VerifyPassword(password, []byte(pw))
	if err != nil {
		return false
	}

	return true
}

func Login(w http.ResponseWriter, r *http.Request) {

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		Logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pw, ok := AdminCreds[creds.Username]
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)

		return
	}

	err := VerifyPassword(creds.Password, []byte(pw))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)

		return
	}

	w.WriteHeader(http.StatusOK)

}

func CreatePasswordHash(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		Logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := HashPassword(creds.Password)
	if err != nil {
		Logger.Error("Error hashing password: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(hash), creds.Username)

	w.WriteHeader(http.StatusOK)

}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
