import { defineStore } from "pinia"
import type { Event, DownTimeType, Location, Credentials } from "../types/types.ts"
import type { ToastMessageOptions } from "primevue/toast"
import LoginView from "@/views/LoginView.vue"


const creds = "Basic " + btoa("mybasicauth:password")
const url = "http://localhost:8080/"

export const useDataStore = defineStore("dataStore", {
  state: () => ({
    events: [

    ] as Event[],
    dateRange: {
      startDate: "2025-01-27T00:00:00.000Z",
      endDate: "2025-01-29T00:00:00.000Z"
    },
    downTimeTypes: new Array as DownTimeType[],
    locations: new Array as Location[],

    credentials: {
      username: "",
      password: ""
    } as Credentials,

    toast: {
      severity: "info",
      summary: "",
      detail: "",
      life: 3000,
    } as ToastMessageOptions,
    refreshInterval: 120,


  }),
  getters: {
    getEvents(state) {
      return state.events
    },
    getDateRange(state) {
      return state.dateRange
    },
    getToast(state) {
      return state.toast
    }


  },
  actions: {

    setDate(start: string, end: string) {
      this.dateRange.startDate = start
      this.dateRange.endDate = end
    },

    async fetchEvents(start: string, end: string) {
      try {
        const response = await fetch(`${url}events?startDate=${start}&endDate=${end}`, {
          headers: {
            "Authorization": creds
          }
        })
        const data = await response.json()
        this.events = data

      } catch (err) {
        this.toast = {
          severity: "error",
          summary: "Eventabfrage fehlgeschlagen",
          detail: "Daten konnten nicht geladen werden",
          life: 5000
        }
      }
    },
    async login(username: string, password: string) {


    },

    async addEvent(event: Event) {
      try {
        const response = await fetch(`${url}events`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": creds
          },
          body: JSON.stringify(event)
        })
        if (response.status >= 400) {
          throw new Error("Bad response from server: " + response.status)
        }


      } catch (error) {
        this.toast = {
          severity: "error",
          summary: "Fehler beim Speichern",
          detail: "Event konnte nicht hinzugefügt werden",
          life: 5000
        }
      }
    },


    async addDownTimeType(dt: DownTimeType) {
      try {
        const response = await fetch(`${url}downtime_types`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": creds
          },
          body: JSON.stringify(dt)
        })
        if (response.status >= 400) {
          throw new Error("Bad response from server: " + response.status)
        }

      } catch (error) {
        this.toast = {
          severity: "error",
          summary: "Fehler beim Speichern",
          detail: "Typ konnte nicht hinzugefügt werden",
          life: 5000
        }
      }
    },
    async addLocation(location: Location) {
      try {
        const response = await fetch(`${url}locations`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": creds
          },
          body: JSON.stringify(location)
        })
        if (response.status >= 400) {
          throw new Error("Bad response from server: " + response.status)
        }

      } catch (error) {
        this.toast = {
          severity: "error",
          summary: "Fehler beim Speichern",
          detail: "Standort konnte nicht hinzugefügt werden",
          life: 5000
        }
      }
    },
    async getDownTimeTypes() {
      try {

        const response = await fetch(`${url}downtime_types`, {
          headers: {
            "Authorization": creds
          }
        })
        const data = await response.json()
        this.downTimeTypes = data
      } catch (error) {
        this.toast = {
          severity: "error",
          summary: "Fehler beim Abruf",
          detail: "Typen konnten nicht abgerufen werden",
          life: 5000
        }
      }
    },
    async getLocations() {
      try {
        const response = await fetch(`${url}locations`, {
          headers: {
            "Authorization": creds
          }
        })
        const data = await response.json()
        this.locations = data
      } catch (error) {
        this.toast = {
          severity: "error",
          summary: "Fehler beim Abruf",
          detail: "Standorte konnten nicht abgerufen werden",
          life: 5000
        }
      }
    },
    async removeEvent(id: string) {
      try {
        const response = await fetch(`${url}events/${id}`, {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
            "Authorization": creds
          },

        })
        if (response.status >= 400) {
          throw new Error("Bad response from server: " + response.status)
        }
        return true
      } catch (error) {
        this.toast = {
          severity: "error",
          summary: "Fehler beim Löschen",
          detail: "Event konnte nicht gelöscht werden",
          life: 5000
        }
        return false
      }


    }

  }
})