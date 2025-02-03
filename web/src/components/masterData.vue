<script setup lang="ts">

import { ref } from "vue";
import { useDataStore } from "@/stores/dataStore";


import InputText from 'primevue/inputtext';
import ColorPicker from 'primevue/colorpicker';
import Button from 'primevue/button';
import type { DownTimeType, Location } from "@/types/types";


const store = useDataStore();

const location = ref({
    name: "",
    color: "",
} as Location)

const downtimeType = ref({
    name: "",
    color: "",
} as DownTimeType)

function CreateLocation() {
    if (location.value.name === "" || location.value.color === "") {
        return;
    }
    store.addLocation(location.value);
}

function CreateDownTimeType() {
    if (downtimeType.value.name === "" || downtimeType.value.color === "") {
        return;
    }
    store.addDownTimeType(downtimeType.value);
}





</script>

<template>
    <section>
        <div class="md-tile">
            <h2>Standort Erstellen</h2>
            <div class="input-group">
                <label>Standortname</label>
                <InputText id="location" v-model="location.name" />
            </div>
            <div class="input-group">
                <label>Farbe</label>
                <ColorPicker v-model="location.color" />
            </div>
            <Button label="Erstellen" style="background-color: var(--color-primary-1);" @click="CreateLocation"/>
        </div>
        <div class="md-tile">
            <h2>Stillstandstypen Anlegen</h2>
            <div class="input-group">
                <label>Typname</label>
                <InputText id="type" v-model="downtimeType.name" />
            </div>
            <div class="input-group">
                <label>Farbe</label>
                <ColorPicker v-model="downtimeType.color" />
            </div>
            <Button label="Erstellen" style="background-color: var(--color-primary-1);" @click="CreateDownTimeType"/>

        </div>
    </section>


</template>

<style>
.md-tile {
    display: flex;
    align-items: flex-start;
    flex-direction: column;
    margin: 10px 2px;
    padding: 1rem;
    border: 1px solid rgb(231, 231, 231);
    border-radius: 2px;
}

.md-tile>h2 {
    font-style: italic;
}

.input-group {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 10px 0;

}
</style>