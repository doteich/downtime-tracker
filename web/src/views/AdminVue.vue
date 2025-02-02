<script setup lang="ts">

import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';


</script>




<template>
    <section>
        <Tabs value="0">
            <TabList>
                <Tab value="0">Erstellen</Tab>
                <Tab value="1">Ändern</Tab>

            </TabList>
            <TabPanels>
                <TabPanel value="0">
                    <p>
                        A
                    </p>
                </TabPanel>
                <TabPanel value="1">
                    <p>
                        B
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