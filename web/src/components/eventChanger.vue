<script setup lang="ts">
import { ref, computed } from "vue"
import { useDataStore } from "@/stores/dataStore"
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from "primevue/button"
import DatePicker from 'primevue/datepicker';


const store = useDataStore()

const end = ref(new Date())
const start = ref(new Date())



const events = computed(() => {
    return store.getEvents
})

function reload() {
    store.fetchEvents(start.value.toISOString(), end.value.toISOString())
}

async function remove(props: any) {
    let res = await store.removeEvent(props.id)
    if (res) {
        reload()
    }

}


</script>

<template>

    <section class="event-changer">
        <div class="ec-header">
            <div>
                <label>Startdatum</label>
                <DatePicker v-model="start" showTime hourFormat="24" fluid dateFormat="dd.mm.yy"></DatePicker>
            </div>
            <div>
                <label>Enddatum</label>
                <DatePicker v-model="end" showTime hourFormat="24" fluid dateFormat="dd.mm.yy"></DatePicker>
            </div>
            <Button icon="bi bi-arrow-clockwise" style="background: var(--color-primary-1);" @click="reload"></Button>

        </div>


        <DataTable :value="events" tableStyle="min-width: 50rem">
            <Column field="id" header="ID"></Column>
            <Column field="name" header="Name"></Column>
            <Column field="startDate" header="Start"></Column>
            <Column field="endDate" header="Ende"></Column>
            <Column field="type" header="Typ"></Column>
            <Column field="location" header="Location"></Column>
            <Column field="actions" header="Aktionen">
                <template #body="slotProps">
                    <Button icon="bi bi-trash-fill" severity="info" @click="remove(slotProps.data)"></Button>
                </template>
            </Column>

        </DataTable>


    </section>
</template>

<style>
.event-changer {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    padding: 1rem;
    border-radius: 4px;


    box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);

}


.ec-header {
    display: flex;
    align-items: flex-end;
    width: 100%;
}

.ec-header>div {
    display: flex;
    flex-direction: column;
    margin: 0 10px;
}
</style>
