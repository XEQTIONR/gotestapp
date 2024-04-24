<template>
  <div class="about">
    <h1>This is an about page</h1>
    <span>{{ watchman }}</span>
    <div :style="{'background-color': color ?? 'grey' }" class="status"></div>
    <div>Specie : {{ specie }}</div>
    <div>Age: {{ age }}</div>
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
<!-- 
<script setup>
  import { onBeforeRouteLeave, onBeforeRouteUpdate } from '@/router/vue-router.mjs'
  import { ref } from 'vue'

  const post = ref(null)
  const age = ref(null)
  const color = ref(null)
  const specie = ref(null)
  const busy = ref(false)

  function setData() {
      busy.value = true
      if (window.originURL != window.location.href || !window.cacheDirty) {
        window.cacheDirty = true
        console.log('making ajax call to ' + window.location.href)
        axios.get(window.location.href, {
          headers: {
            'AJAXRequest' : 'true',
          }
        })
        .then((res) => {
          console.log('I am ajax -- response data:', res.data)

          const { age, color, specie } = res.data.data
          age.value = age
          color.value = color
          specie.value = specie
          busy.value = false
        })
      } else {
        console.log('not making API Call')
        console.log(window.apiData)
        
        const { age, color, specie } = window.apiData
        age.value = age
          color.value = color
          specie.value = specie
          busy.value = false
      }
  }

  function beforeRouteEnter(to, from, next) {
    console.log('AboutView beforeRouteEnter')
  }
  
  onBeforeRouteLeave((to, from) => {

    console.log('AboutView onBeforeRouteLeave')
    console.log('to: ', to)
    console.log('from: ', from)
    console.log('post', post)
    console.log('post.value', post.value)
  }) 

  onBeforeRouteUpdate((to, from) => {
    console.log('AboutView onBeforeRouteUpdate')
  } )

</script> -->

<script>
  export default {
  data () {
    return {
      post: null,
      age: null,
      color: null,
      specie: null,
      busy: false,
    }
  },
  beforeRouteEnter (to, from, next) {
      console.log('AboutView beforeRouteEnter')
      next(vm => vm.setData())
  },
  // when route changes and this component is already rendered,
  // the logic will be slightly different.
  beforeRouteUpdate (to, from, next) {
    console.log('AboutView beforeRouteUpdate')
      next()
  },

  methods: {
    setData() {
      this.busy = true

      if (window.originURL == window.location.href && window.valid) {
        console.log('window.originURL == window.location.href && window.valid')
        console.log('not making API Call')
        console.log(window.apiData)
        
        const { age, color, specie } = window.apiData
        this.age = age
        this.color = color
        this.specie = specie
        this.busy = false
        window.valid = false
      } else { // ajax page load 
        console.log('making ajax call to ' + window.location.href)
        axios.get(window.location.href, {
          headers: {
            'AJAXRequest' : 'true',
          }
        })
        .then((res) => {
          console.log('I am ajax -- response data:', res.data)

          const { age, color, specie } = res.data.data
          this.age = age
          this.color = color
          this.specie = specie
          this.busy = false
          window.valid = false
        })
      }
    }
  },

  computed: {
    postJson() {
      return JSON.stringify(this.post)
    },

    statusColor() {
      if (this.busy) {
        return 'red';
      }

      return 'green';
    }
  }
}
</script>