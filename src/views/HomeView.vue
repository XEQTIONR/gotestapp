<script setup>
import TheWelcome from '../components/TheWelcome.vue'
//import { RouterLink } from 'vue-router'
import { RouterLink } from '@/router/vue-router.mjs'
</script>

<template>
  <main>
    <TheWelcome />
    <RouterLink to="/about">To About</RouterLink>
  </main>
</template>
