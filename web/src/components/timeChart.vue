<script setup lang="ts">
import { computed, ref } from 'vue';
import { useDataStore } from "@/stores/dataStore";
import Chart from 'chart.js/auto';
import ChartDataLabels from 'chartjs-plugin-datalabels';


Chart.register(ChartDataLabels);

const store = useDataStore();



const chartOptions = computed(() => {

  let range = store.getDateRange;
  let locations = store.locations;
  console.log(locations)
  let loc = locations?.map(l => l.name);


  return {
    indexAxis: 'y',
    maintainAspectRatio: false,
    aspectRatio: 1,
    responsive: true,
    scales: {
      x: {
        type: 'time',
        time: {
          unit: 'hour', // Major ticks for days
          displayFormats: {
            day: 'MMM dd', // Format for major ticks (e.g., "Jan 27")
            hour: 'HH:mm', // Format for minor ticks (e.g., "09:00")
          },
        },
        ticks: {
          major: {
            enabled: true, // Enable major ticks for days
          },

        },
        min: range.startDate, // Start of the week
        max: range.endDate, // End of the week
      },
      y: {
        type: 'category',
        labels: loc, // Y-axis labels
      },
    },

    plugins: {
      datalabels: {
        color: 'black',
        borderRadius: 2,
        borderColor: "grey",

        align: "start",
        formatter: function (value: string, context: any) {
          console.log(value, context);
          return context.dataset.label + "➡️"
        }
      }
    }
  }


})




const chartPlugins = [ChartDataLabels]



const dat = computed(() => {
  const events = store.getEvents;
  let datasets = events.map(event => ({
    label: event.name,
    data: [{
      x: [event.startDate, event.endDate], // Time range for the bar
      y: event.location, // Location on the y-axis
    }],
    backgroundColor: "#"+event.color,

  }));

  return {
    datasets: datasets, // Wrap datasets in a `datasets` object
  };
});




</script>


<template>
  <div style="height: 80vh">
    <de-chart type="bar" :data="dat" :options="chartOptions" :plugins="chartPlugins"></de-chart>
  </div>
</template>


<style>
canvas {
  height: 80vh !important;
  color: rgba(126, 126, 126, 0.527);
}
</style>