<script setup lang="ts">
import { ref } from 'vue'
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';

const username = ref('');
const password = ref('');


</script>


<template>
    <section class="login-mask">

        <div>
            <h1>Login</h1>
            <InputText v-model="username" placeholder="Username" />
            <InputText v-model="password" placeholder="Password" type="password" />
            <Button label="Login" @click="" style="background: var(--color-primary-1);" />
        </div>
    </section>

</template>

<style>
.login-mask {
    display: flex;
    height: 50vh;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 0 1rem;
}

.login-mask>div {

    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
    margin: 1rem;
    max-width: 600px;
   width: 100%;
    border-radius: 5px;
    background-color: #f5f5f5;
}
.login-mask>div > h1 {
   
    width: 100%;
    color: white;
    color: var(--color-primary-1);

}

</style>