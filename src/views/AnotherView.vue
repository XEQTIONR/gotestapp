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
      pageData: null,
      busy: false,
    }
  },

  methods: {
    setData() {
      this.busy = true
      useData()
        .then(response => {
          this.pageData = response
          this.busy = false
        })
    }
  },
}


</script>