<script setup lang="ts">

import { ref } from "vue";
import { useDataStore } from "@/stores/dataStore";
import InputText from 'primevue/inputtext';
import ColorPicker from 'primevue/colorpicker';
import Button from 'primevue/button';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';



import type { DownTimeType, Location } from "@/types/types";


const store = useDataStore();

const location = ref({
    name: "",
    color: "2e2b6e",
} as Location)

const downtimeType = ref({
    name: "",
    color: "2e2b6e",
} as DownTimeType)

function CreateLocation() {
    
    if (location.value.name === "" || location.value.color === "") {
        return;
    }
    
    
    store.addLocation(location.value);
    location.value.name = ""
}

function CreateDownTimeType() {
    if (downtimeType.value.name === "" || downtimeType.value.color === "") {
        return;
    }

   
    store.addDownTimeType(downtimeType.value);
    downtimeType.value.name = ""
}





</script>

<template>
    <section class="master-data">
        <div class="md-tile">
            <h2>Standort Erstellen</h2>
  
            <label>(Farbe) Standortname</label>
            <InputGroup>

                <InputGroupAddon>
                    <ColorPicker v-model="location.color" />
                </InputGroupAddon>
                <InputText id="location" v-model="location.name" />
                <Button label="Erstellen" style="background-color: var(--color-primary-1);" @click="CreateLocation"
                    :disabled="location.name == ''" />
            </InputGroup>

        </div>

        <div class="md-tile">
            <h2>Stillstandstypen Anlegen</h2>
            <label>(Farbe) Typname</label>
            <InputGroup>

                <InputGroupAddon>
                    <ColorPicker v-model="downtimeType.color" />
                </InputGroupAddon>
                <InputText id="type" v-model="downtimeType.name" />
                <Button label="Erstellen" style="background-color: var(--color-primary-1);" @click="CreateDownTimeType"
                    :disabled="downtimeType.name == ''" />
            </InputGroup>

        </div>

    </section>


</template>

<style>
.master-data {
    display: flex;
    flex-direction: column;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;

}


.md-tile {
    display: flex;
    align-items: flex-start;
    flex-direction: column;
    margin: 10px 2px;
    padding: 1rem;
    min-width: 60%;
    max-width: 500px;
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