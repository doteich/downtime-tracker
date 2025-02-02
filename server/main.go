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

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool                                 *pgxpool.Pool
	Logger                                 *slog.Logger
	pgHost, pgUser, pgPassword, pgDatabase string
)

type Event struct {
	Name     string `json:"name"`
	Start    string `json:"startDate"`
	End      string `json:"endDate"`
	Location string `json:"location"`
	Color    string `json:"color"`
}

func init() {

	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	flag.StringVar(&pgHost, "host", "localhost", "Postgres host")
	flag.StringVar(&pgUser, "user", "postgres", "Postgres user")
	flag.StringVar(&pgPassword, "password", "pass", "Postgres password")
	flag.StringVar(&pgDatabase, "db", "downtime", "Postgres database")

	flag.Parse()

}

func main() {
	// Initialize the database connection pool
	if err := initDB(); err != nil {
		Logger.Error("Error initializing database connection pool: " + err.Error())
		return
	}
	defer dbPool.Close()

	// Create a new Chi router
	r := chi.NewRouter()

	r.Post("/events", CreateEvent)
	r.Get("/events", GetEvents)

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
	return nil
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		Logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := dbPool.Exec(context.Background(), "INSERT INTO events (name, start_date, end_date, location, color) VALUES ($1, $2, $3, $4, $5)", event.Name, event.Start, event.End, event.Location, event.Color)
	if err != nil {
		Logger.Error("Error inserting event: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	rows, err := dbPool.Query(context.Background(), "SELECT name, start_date, end_date, location, color FROM events")
	if err != nil {
		Logger.Error("Error querying events: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	events := make([]Event, 0)
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.Name, &event.Start, &event.End, &event.Location, &event.Color); err != nil {
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
