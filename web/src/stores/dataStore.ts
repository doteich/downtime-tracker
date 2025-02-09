import { defineStore } from "pinia"
import type { Event, DownTimeType, Location } from "../types/types.ts"
import type { ToastMessageOptions } from "primevue/toast"


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

    toast: {
      severity: "info",
      summary: "",
      detail: "",
      life: 3000,
    } as ToastMessageOptions

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
        const response = await fetch(`${url}events?startDate=${start}&endDate=${end}`)
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
    async addEvent(event: Event) {
      try {
        const response = await fetch(`${url}events`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
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
            "Content-Type": "application/json"
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
            "Content-Type": "application/json"
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

        const response = await fetch(`${url}downtime_types`)
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
        const response = await fetch(`${url}locations`)
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
    }

  }
})