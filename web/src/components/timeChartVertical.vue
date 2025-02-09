<script setup lang="ts">
import { computed, ref } from 'vue';
import { useDataStore } from "@/stores/dataStore";
import Chart from 'chart.js/auto';
import ChartDataLabels from 'chartjs-plugin-datalabels';

Chart.register(ChartDataLabels);

const store = useDataStore();

const chartOptions = computed(() => {
  return {
    indexAxis: 'x',
    maintainAspectRatio: false,
    aspectRatio: 1,
    responsive: true,
    scales: {
      x: {
        type: 'category',
        labels: ['Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag'],
      },
      y: {
        type: 'linear',
        min: 0,
        reverse: true,
        max: 24,
        ticks: {

          stepSize: 1,
          callback: function (value: number) {
            return `${value}:00`;
          }
        }
      },
    },

    plugins: {

      legend: {
        labels: {
          generateLabels: function (chart: any) {
            let labels = [] as any[]
            chart.data.datasets.forEach((el: any) => {
              if (!labels.find(x => x.text === el.label)) {
                labels.push({
                  text: el.label,
                  fillStyle: el.backgroundColor,

                })
              }
            });
            return labels
          }
        }
      },
      datalabels: {
        color: 'black',
        borderRadius: 2,
        borderColor: "black",
        backgroundColor: "lightgrey",
        align: "center",
        formatter: function (value: string, context: any) {
         
          return context.dataset.data[0].name;
        }
      }
    }
  };
});

const chartPlugins = [ChartDataLabels];

const dat = computed(() => {
  let events = store.getEvents;
  let datasets = events.map(event => {
    const startDate = new Date(event.startDate);
    const endDate = new Date(event.endDate);

    const dayStart = getWeekday(startDate);
    const dayEnd = getWeekday(endDate);
    const startHour = startDate.getHours() + startDate.getMinutes() / 60;
    const endHour = endDate.getHours() + endDate.getMinutes() / 60;

    if (dayStart !== dayEnd) {
      return {
        label: event.location,
        data: [{
          x: dayStart,
          y: [startHour, 24],
          name: event.name,
        },
        {
          x: dayEnd,
          y: [0, endHour],
          name: event.name,
        }],
        backgroundColor: setLocationColor(event.location),
      };
    }

    return {
      label: event.location,
      data: [{
        x: dayStart,
        y: [startHour, endHour],
        name: event.name,
      }],
      backgroundColor: setLocationColor(event.location),
    };
  });

  return {

    datasets: datasets,
  };
});

const getWeekday = (date: Date) => {
  const days = ['Sonntag', 'Montag', 'Dienstag', 'Mittwoch', 'Donnerstag', 'Freitag', 'Samstag'];
  return days[date.getDay()];
};

const setLocationColor = (l: string) => {

  const location = store.locations.find(x => x.name === l);


  return location ? "#" + location.color : 'green';

}



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