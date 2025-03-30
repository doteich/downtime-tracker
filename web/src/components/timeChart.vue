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

function setShift(d: Date, label: string, start: number, end: number, height: number) {

  let dayStart = new Date(d.setHours(start, 0, 0, 0));
  let dayEnd = new Date(d.setHours(end, 0, 0, 0));

  if (start > end) {
    dayEnd.setDate(dayEnd.getDate() + 1)
  }


  return {
    type: 'box',
    xMin: dayStart.getTime(),
    xMax: dayEnd.getTime(),
    yMin: -0.5,
    yMax: height - 0.5,
    backgroundColor: 'rgba(0, 0, 0, 0.00)',
    borderColor: 'rgba(58, 64, 74,0.4)',
    borderWidth: 1,
    borderDash: [3, 3],
    label: {
      content: label,

      display: true,
      position: {
        x: 'center',
        y: 'end'
      },
    }
  }
}

const chartOptions = computed(() => {

  let range = store.getDateRange;

  let locations = store.locations;
  let loc = locations?.map(l => l.name);

  let days = eachDayOfInterval({
    start: new Date(range.startDate),
    end: new Date(range.endDate)
  });


  let boxes = {} as any;

  days.forEach((day, index) => {

    let d = new Date(day);
    let dayName = getWeekday(d.getDay());

    let dayStart = new Date(d.setHours(0, 0, 0, 0));
    let dayEnd = new Date(d.setHours(23, 59, 59, 999));

    if (dayName == "Sonntag") {
      dayStart = new Date(d.setHours(18, 0, 0, 0));
      dayName = "Son"
    }

    console.log(dayName, dayStart, dayEnd)

    boxes[dayName + index] = {
      type: 'box',
      xMin: dayStart.getTime(),
      xMax: dayEnd.getTime(),
      yMin: -0.5,
      yMax: locations.length - 0.5,
      backgroundColor: 'rgba(0, 0, 0, 0.00)',
      label: {
        content: dayName,
        font: {
          size: 18
        },

        display: true,
        position: {
          x: 'center',
          y: 'start'
        },
      }
    }

    boxes[dayName + "fs"] = setShift(d, "FrS", 6, 14, locations.length)
    boxes[dayName + "ss"] = setShift(d, "SpS", 14, 22, locations.length)
    boxes[dayName + index + "ns"] = setShift(d, "NaS", 22, 6, locations.length)
  });

  boxes["runningLine"] = {
    type: 'line',
    xMin: new Date(),
    xMax: new Date(),
    borderColor: 'rgb(255, 99, 132)',
    borderWidth: 2,
    borderDash: [3, 3],
  }


  let start = new Date(range.startDate)

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
        min: start, // Start of the week
        max: range.endDate, // End of the week
      },
      y: {
        offset: true,
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

  let locMap = new Map();
  events.forEach(event => {
    let key = event.location;
    if (!locMap.has(key)) {
      locMap.set(key, []);
    }

  });


  let datasets = events.map(event => {

    let el = locMap.get(event.location)
    let index = el.length;
    el.push(event.name)

    let stack = `stack${index}`


    return {
      label: event.name,
      data: [{
        x: [event.startDate, event.endDate],
        y: event.location,

      }],
      barPercentage: 4,
      maxBarThickness: 50,
      stack: stack,
      backgroundColor: "#" + event.color,

    }
  })
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