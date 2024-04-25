<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div>Specie : {{ all?.specie }}</div>
    <div>Age: {{ all?.age }}</div>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }

  .status {
    width: 50px;
    height: 50px;
  }
}
</style>

<script>
import { useData } from '../composables/useData.js'

export default {
  data () {
    return {
      all: null,
      busy: false,
    }
  },
  beforeRouteEnter (to, from, next) {
    next(vm => vm.setData())
  },
  // when route changes and this component is already rendered,
  // the logic will be slightly different.
  beforeRouteUpdate (to, from, next) {
    next()
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