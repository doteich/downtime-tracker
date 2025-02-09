import './assets/main.css'
import "bootstrap-icons/font/bootstrap-icons.css";

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config';
import Chart from 'primevue/chart';

import Aura from '@primevue/themes/aura';
import 'chartjs-adapter-date-fns';


import ToastService from 'primevue/toastservice';

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
    // Default theme configuration
    theme: {
        preset: Aura,
        options: {
            prefix: 'p',
            darkModeSelector: 'off',
            
        }
    }
 });

 app.use(ToastService);

 app.component('de-chart', Chart);



app.mount('#app')
