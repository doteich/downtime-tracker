<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { RouterLink, RouterView } from 'vue-router';
import router from './router';
import { useDataStore } from './stores/dataStore';

const store = useDataStore();

// State to track whether the full header is shown
const showFullHeader = ref(true);

// Function to toggle the header view
const toggleHeader = () => {
  showFullHeader.value = !showFullHeader.value;
};



onMounted(() => {
  store.getDownTimeTypes();
  store.getLocations();
  store.fetchEvents();
});



</script>

<template>
  <div class="app-container">
    <header :class="{ overlay: !showFullHeader }">
      <h1 v-if="showFullHeader" style="font-style: italic;">{{ router.currentRoute.value.meta.title }}</h1>
      <nav>
        <ul class="nav-list">
          <li>
            <!-- Button to toggle the header -->
            <button @click="toggleHeader"><i class="bi bi-box-arrow-in-up-right"></i></button>
          </li>
          <li>
            <RouterLink to="/"><i class="bi bi-house-door"></i></RouterLink>
          </li>
          <li>
            <RouterLink to="/admin"><i class="bi bi-wrench-adjustable"></i></RouterLink>
          </li>
        </ul>
      </nav>

      <!-- Conditionally render the full header content -->

    </header>


    <main>
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
</style>
