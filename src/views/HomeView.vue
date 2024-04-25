<script>
import TheWelcome from '../components/TheWelcome.vue'

import { useData } from '../composables/useData.js'

export default {
  
  components: {
    TheWelcome,
  },

  data() {
    return {
      busy: false,
      all: null,
    }
  },
  beforeRouteEnter (to, from, next) {
    next(vm => {
      vm.setData()
  })
  },

  methods: {
    setData() {
      useData()
        .then(response => {
          this.all = response
        })
    }
  }

}
</script>

<template>
  <main>
    <TheWelcome />
    <ol v-if="all">
      <li v-for="person in all.people">{{ person.name }}</li>
    </ol>
  </main>
</template>
