<script>
import TheWelcome from '../components/TheWelcome.vue'
//import { RouterLink } from 'vue-router'
// import { RouterLink } from '@/router/vue-router.mjs'

export default {
  
  components: {
    TheWelcome,
  },

  data() {
    return null
  },
  beforeRouteEnter (to, from, next) {
    console.log('HomeView beforeRouteEnter')
    console.log('window.originURL:', window.originURL)
      console.log('window.apiData:' , window.apiData)
    next()
  }

}
</script>

<template>
  <main>
    <TheWelcome />
  </main>
</template>
