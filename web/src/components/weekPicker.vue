<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Button from 'primevue/button';
import { getWeek, setWeek, lastDayOfWeek, startOfWeek, addSeconds } from "date-fns"
import { useDataStore } from "@/stores/dataStore"

const store = useDataStore();

const week = ref(0);
const week_number = ref(0);
const week_start = ref("");
const week_end = ref("");


onMounted(() => {
    let date = new Date();
    let w = getWeek(date);
    week.value = w
    week_number.value = w;
    setWeekDates(w);

    setInterval(() => { refresh() }, store.refreshInterval * 1000)

});

const refresh = () => {
    let date = new Date();
    let w = getWeek(date);
    week.value = w
    week_number.value = w;
    setWeekDates(w);
}


const updateWeek = (delta: number) => {
    if (week.value == 1 && delta == -1) {
        week.value = 53;
    } else if (week.value == 52 && delta == 1) {
        week.value = 0;
    }
    week.value += delta;
    week_number.value += delta;
    setWeekDates(week_number.value);
}

const setWeekDates = (w: number) => {
    let d = setWeek(new Date(), w);
    let start = startOfWeek(d, { weekStartsOn: 0 })
    start.setHours(18, 0, 0, 0)

    if (d.getDay() < 2) {
        d.setHours(d.getHours() + 24)
    }

    let end = lastDayOfWeek(d, { weekStartsOn: 0 })
    end.setDate(end.getDate() + 1)

    store.setDate(start.toISOString(), end.toISOString());

    week_start.value = start.toLocaleDateString();
    week_end.value = end.toLocaleDateString();

    store.fetchEvents(start.toISOString(), end.toISOString());
}


</script>


<template>
    <div class="week-picker">
        <Button icon="bi bi-chevron-double-left" rounded severity="info" aria-label="Left"
            style="background-color: var(--color-primary-1)" @click="updateWeek(-1)" />

        <div>
            <p class="week-main">KW <span class="week">{{ week }}</span></p>
            <p class="week-range">{{ week_start }}-{{ week_end }}</p>
        </div>

        <Button icon="bi bi-chevron-double-right" severity="info" rounded aria-label="Right"
            style="background-color: var(--color-primary-1)" @click="updateWeek(1)" />
    </div>


</template>


<style>
.week-picker {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 0.2rem;
    background-color: #f0f0f0;
    border-radius: 5px;
    margin: 1rem;
}

.week-picker>div {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.week-main {
    font-size: 1.5rem;
    font-weight: bold;
    color: var(--color-primary-1);
}

.week {
    font-size: 1.7rem;
    font-weight: bold;
    font-style: italic;
    color: var(--color-primary-1);
}

.week-range {
    font-size: 0.9rem;
    font-style: italic;
    color: grey;
}
</style>
