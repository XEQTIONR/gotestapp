<template>
    AnotherView.vue
</template>

<script>
import { useData } from '../composables/useData.js'

export default {
  beforeRouteEnter (to, from, next) {
    next(vm => vm.setData())
  },

  data () {
    return {
      all: null,
      busy: false,
    }
  },

  methods: {
    setData() {
      this.busy = true
      useData()
        .then(response => {
          this.all = response
        })
    }
  },
}


</script>