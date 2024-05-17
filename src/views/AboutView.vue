<template>
  <AppHeader />
  <div class="about">
    <h1>This is an about page</h1>
    
      <dl>
        <dt>Color</dt>
        <dd><div :style="{ width: '50px', height: '50px', 'background-color': pageData?.color ?? 'black' }"></div></dd>
        <dt>Specie</dt>
        <dd>{{ pageData?.specie }}</dd>
        <dt>Age</dt>
        <dd>{{ pageData?.age }}</dd>
      </dl>
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

import AppHeader from '../components/Header.vue'

export default {

  components: {
    AppHeader
  },
  data () {
    return {
      pageData: null,
    }
  },
  beforeRouteEnter (to, from, next) {
    next(vm => vm.pageData = window.data)
  },
}
</script>