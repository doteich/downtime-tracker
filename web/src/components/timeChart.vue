<script setup lang="ts">
import { computed, ref } from 'vue';
import { useDataStore } from "@/stores/dataStore";
import Chart from 'chart.js/auto';
import ChartDataLabels from 'chartjs-plugin-datalabels';
import annotationPlugin from 'chartjs-plugin-annotation';
import { eachDayOfInterval } from "date-fns";

Chart.register(ChartDataLabels);
Chart.register(annotationPlugin);

const store = useDataStore();

const getWeekday = (day: number) => {
  const days = ['Sonntag', 'Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag', 'Samstag'];
  return days[day];
};


const chartOptions = computed(() => {

  let range = store.getDateRange;

  let locations = store.locations;
  let loc = locations?.map(l => l.name);

  let days = eachDayOfInterval({
    start: new Date(range.startDate),
    end: new Date(range.endDate)
  });



  let boxes = {} as any;

  days.forEach(day => {
    let d = new Date(day);
    let dayName = getWeekday(d.getDay());
    let dayStart = new Date(d.setHours(0, 0, 0, 0));
    let dayEnd = new Date(d.setHours(23, 59, 59, 999));
    boxes[dayName] = {
      type: 'box',
      xMin: dayStart.getTime(),
      xMax: dayEnd.getTime(),
      yMin: -0.5,
      yMax: locations.length - 0.5,
      backgroundColor: 'rgba(0, 0, 0, 0.00)',
      label: {
        content: dayName,
        display: true,
        position: {
          x: 'center',
          y: 'end'
        },
      }
    }
  });

  console.log(boxes)



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
      tooltip: {
        callbacks: {
          label: function (context: any) {
            let d_start = new Date(context.parsed._custom.barStart).toLocaleString()
            let d_end = new Date(context.parsed._custom.barEnd).toLocaleString()
            return `${d_start} - ${d_end}`
          },
          title: function (context: any) {
            return context[0].dataset.label
          }
        }
      },
      datalabels: {
        color: 'black',
        borderRadius: 4,
        borderColor: "grey",
        backgroundColor: "rgba(230, 230, 230, 0.844)",
        align: "center",
        formatter: function (value: string, context: any) {

          return context.dataset.label
        }
      },
      annotation: {
        annotations: boxes

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
    barPercentage: 2,
    maxBarThickness: 50,
    backgroundColor: "#" + event.color,

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