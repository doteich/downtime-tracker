<script setup lang="ts">
import { ref, onMounted } from "vue";

import { useDataStore } from "@/stores/dataStore";


const store = useDataStore();

onMounted(() => {
    store.getDownTimeTypes();
    store.getLocations();
});



</script>

<template>
    <div>
        {{ store.locations }}
        {{ store.downTimeTypes }}

    </div>


</template>

<style scoped></style>