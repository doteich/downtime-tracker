<script setup lang="ts">
import { onMounted } from 'vue';
import { useDataStore } from '@/stores/dataStore';
import router from '@/router';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';
import masterData from '../components/masterData.vue';
import eventCreator from '@/components/eventCreator.vue';
import eventChanger from '@/components/eventChanger.vue';

const store = useDataStore();

onMounted(() => {
    if (store.credentials.username === '' || store.credentials.password === '') {
        router.push('/login');
    }
});



</script>


<template>
    <section>
        <Tabs value="0">
            <TabList>
                <Tab value="0">Erstellen</Tab>
                <Tab value="1">Ändern</Tab>
                <Tab value="2">Stammdaten</Tab>
            </TabList>
            <TabPanels>
                <TabPanel value="0">
                    <p>
                        <event-creator></event-creator>
                    </p>
                </TabPanel>
                <TabPanel value="1">
                    <p>
                        <event-changer></event-changer>
                    </p>
                </TabPanel>
                <TabPanel value="2">
                    <p>
                        <master-data></master-data>
                    </p>
                </TabPanel>
            </TabPanels>
        </Tabs>

    </section>


</template>

<style>
.p-tab-active {
    color: var(--color-primary-1) !important;

}

.p-tablist-active-bar {
    background-color: var(--color-primary-1) !important;

}
</style>