<script>
import TheWelcome from '../components/TheWelcome.vue'

import { useData } from '../composables/useData.js'

export default {
  
  components: {
    TheWelcome,
  },

  data() {
    return {
      pageData: null,
    }
  },
  beforeRouteEnter (to, from, next) {
    next(vm => vm.pageData = window.data)
  },

}
</script>

<template>
  <main>
    <TheWelcome />
    <ol v-if="pageData">
      <li v-for="person in pageData.people">{{ person.name }}</li>
    </ol>
  </main>
</template>
