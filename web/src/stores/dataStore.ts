import { defineStore } from "pinia"
import type { Event } from "../types/types.ts"

const url = "http://localhost:8080/"


export const useDataStore = defineStore("dataStore", {
  state: () => ({
    events: [
      { location: "Location A", startDate: "2025-01-27T09:00:00.000Z", endDate: "2025-01-27T12:00:00.000Z", name: "Event 1", color: "red" },
      { location: "Location A", startDate: "2025-01-27T09:00:00.000Z", endDate: "2025-01-27T12:00:00.000Z", name: "Event 2", color: "blue" },
      { location: "Location B", startDate: "2025-01-27T14:00:00.000Z", endDate: "2025-01-27T16:00:00.000Z", name: "Event 3", color: "green" },
      { location: "Location C", startDate: "2025-01-28T10:00:00.000Z", endDate: "2025-01-28T15:00:00.000Z", name: "Event 4", color: "lime" },
      { location: "Location D", startDate: "2025-01-29T08:00:00.000Z", endDate: "2025-01-29T10:00:00.000Z", name: "Event 5", color: "gold" },
      { location: "Location D", startDate: "2025-01-30T08:00:00.000Z", endDate: "2025-01-30T10:00:00.000Z", name: "Event 6", color: "violet" },
    ] as Event[],
    dateRange: {
      startDate: "2025-01-27T00:00:00.000Z",
      endDate: "2025-01-29T00:00:00.000Z"
    },

  }),
  getters: {
    getEvents(state) {
      return state.events
    },
    getDateRange(state) {
      return state.dateRange
    }


  },
  actions: {

    setDate(start: string, end: string) {
      this.dateRange.startDate = start
      this.dateRange.endDate = end
    },

    async fetchEvents(start: string, end: string) {
      const response = await fetch(`${url}events`)
      const data = await response.json()
      this.events = data
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
        console.error(error)
      }
    }

  }
})