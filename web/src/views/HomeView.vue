<script setup lang="ts">;
import { ref } from "vue";

import weekPicker from "../components/weekPicker.vue";
import timeChart from "../components/timeChart.vue";
import timeChartVertical from "../components/timeChartVertical.vue";

const chartType = ["bi bi-bar-chart-steps", "bi bi-bar-chart"];
const selection = ref(0);

function toggleChart() {
    selection.value = selection.value === 0 ? 1 : 0;
}

</script>

<template>
    <section>
        <week-picker></week-picker>
        <div class="chart-toggler">
            <button @click="toggleChart()"><i :class="chartType[selection]"></i></button>
        </div>
        <time-chart v-if="selection == 0"></time-chart>
        <time-chart-vertical v-else></time-chart-vertical>
    </section>
</template>

<style>
.chart-toggler {
    display: flex;
    justify-content: flex-end;
    margin: 0rem;
    padding: 0 1rem;

}

.chart-toggler>button {
    margin: 0rem;
    padding: 1px 10px;
    cursor: pointer;
    border: none;
    background-color: transparent;
}
</style>