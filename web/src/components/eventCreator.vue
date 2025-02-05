<script setup lang="ts">
import { ref, computed } from "vue";
import DatePicker from 'primevue/datepicker';
import InputText from 'primevue/inputtext';
import ColorPicker from 'primevue/colorpicker';
import Button from 'primevue/button';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import Select from 'primevue/select';

import { useDataStore } from "@/stores/dataStore";
import type { Event } from "@/types/types";


const store = useDataStore();

const event = ref<Event>({
    name: "",
    location: "",
    type: "",
    startDate: new Date(),
    endDate: new Date(),
    color: "#000000",

});

const locations = computed(() => store.locations);
const downTimeTypes = computed(() => store.downTimeTypes);



</script>

<template>
    <section class="event-creator-container">
        <div class="event-creator">
            <h2>Stillstand anlegen</h2>
            <div class="ei-tile">
                <label>(Farbe) Eventname</label>
                <InputGroup>
                    <InputGroupAddon>
                        <ColorPicker v-model="event.color" />
                    </InputGroupAddon>
                    <InputText id="location" v-model="event.name" placeholder="Stillstand ABC" />
                </InputGroup>
            </div>
            <div class="date-pick">
                <div>
                    <label>Startzeit</label>
                    <DatePicker id="datepicker-24h" v-model="event.startDate" showTime hourFormat="24" fluid
                        dateFormat="dd.mm.yy" />
                </div>
                <div>
                    <label>Endzeit</label>
                    <DatePicker id="datepicker-24h" v-model="event.endDate" showTime hourFormat="24" fluid
                        dateFormat="dd.mm.yy" />
                </div>

            </div>
            <div class="ei-tile">
                <label>Standort</label>
                <Select v-model="event.location" :options="locations" optionLabel="name" placeholder="Standort" />
            </div>

            <div class="ei-tile">
                <label>Stillstandstyp</label>
                <Select v-model="event.type" :options="downTimeTypes" optionLabel="name" placeholder="Typ" />
            </div>


            <div class="ei-tile">
                <Button label="Erstellen" style="background-color: var(--color-primary-1);" />
            </div>
        </div>
    </section>

</template>

<style>

.event-creator-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;

}
.event-creator {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    padding: 1rem;
    border-radius: 4px;
    min-width: 300px;
    max-width: 600px;
    width: 50%;
    box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);
}

.event-creator>h2 {
    background-color: var(--color-primary-1);
    
 
    color: white;
    padding: 0 10px;
    font-style: italic;
}


.ei-tile {
    display: flex;
    flex-direction: column;
    margin: 1rem 0;

}

.date-pick {
    display: flex;
    gap: 1rem;
    margin: 1rem 0;
 
}

.date-pick>div {

    width: 50%;
}
</style>