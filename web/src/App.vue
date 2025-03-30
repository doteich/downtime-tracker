<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { RouterLink, RouterView } from 'vue-router';
import router from './router';
import { useDataStore } from './stores/dataStore';
import { storeToRefs } from 'pinia'
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';



const store = useDataStore();
const toast = useToast();

const { getToast } = storeToRefs(store)

// State to track whether the full header is shown
const showFullHeader = ref(true);

// Function to toggle the header view
const toggleHeader = () => {
  showFullHeader.value = !showFullHeader.value;
};

const showInvModal = ref(false)


onMounted(() => {
  store.getDownTimeTypes();
  store.getLocations();
  store.fetchEvents;
});

watch(getToast, (value) => {

  if (value) {
    toast.add({ severity: value.severity, summary: value.summary, detail: value.detail, life: value.life });
  }
});

function setInterval(i: number) {
  store.refreshInterval = i
}





</script>

<template>
  <div class="app-container">
    <header :class="{ overlay: !showFullHeader }">
      <h1 v-if="showFullHeader" style="font-style: italic;">{{ router.currentRoute.value.meta.title }}</h1>
      <nav>
        <ul class="nav-list">

          <li>

            <button @click="toggleHeader"><i class="bi bi-box-arrow-in-up-right"></i></button>
          </li>
          <li>

            <button @click="showInvModal = !showInvModal"><i class="bi bi bi-clock-history"></i></button>
          </li>

          <li>
            <RouterLink to="/"><i class="bi bi-house-door"></i></RouterLink>
          </li>
          <li>
            <RouterLink to="/login"><i class="bi bi-wrench-adjustable"></i></RouterLink>
          </li>
        </ul>
      </nav>



    </header>


    <main>
      <Dialog v-model:visible="showInvModal" modal header="Aktualisierungsrate (s)" :style="{ width: '25rem' }">
        <div class="modal">
          <input type="number" :value="store.refreshInterval">
          <button>Ok</button>
        </div>
      </Dialog>


      <Toast />
      <RouterView />
    </main>
  </div>
</template>

<style>
/* General app container */
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  position: relative;
}

/* Header styles */
header {
  width: 100%;

  position: fixed;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  top: 0;
  left: 0;
  right: 0;
  z-index: 10;
  border-bottom: 1px solid #e6e6e6;
  transition: all 0.3s ease;
  padding: 10px 20px 0;
}

header.overlay {
  padding: 5px 10px;
  border-bottom: none;
}

.nav-list {
  list-style-type: none;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 5px;
  margin: 0;
  width: fit-content;
  border-radius: 20px;
  background: var(--color-primary-1);
}

.nav-list li {
  margin-right: 10px;
  width: 25px;
  height: 25px;
  text-align: center;


}

.nav-list li>button {
  color: white;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 1.1em;
}

a>i {
  color: white;
}

.nav-list>li:hover>a>i {
  color: gold;
}

.nav-list>li:hover>button>i {
  color: gold;
}

.nav-list li:last-child {
  margin-right: 0;
}

/* Header title styles */
h1 {
  margin-right: auto;
  font-size: 1.5rem;
  color: var(--color-primary-1);

}

/* Main content styles */
main {
  padding: 20px;
  padding-top: 70px;
  /* Space for the full header when visible */
  transition: padding-top 0.3s ease;
}

header.overlay+main {
  padding-top: 10px;
  /* Aligns content to the nav height when the header title is hidden */
}

.modal {
  padding: 5px;
  display: flex;
}
.modal > button{
  background: var(--color-primary-1);
  border: none;
  color: white;
  font-weight: bolder;
  padding:  5px 10px;
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
  cursor: pointer;
}

.modal > button:active{
  color: gold;
}

.modal > input{
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}

</style>
