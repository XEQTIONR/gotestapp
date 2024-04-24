<script>
import TheWelcome from '../components/TheWelcome.vue'
//import { RouterLink } from 'vue-router'
// import { RouterLink } from '@/router/vue-router.mjs'

export default {
  
  components: {
    TheWelcome,
  },

  data() {
    return {
      people: [],
      busy: false,
    }
  },
  beforeRouteEnter (to, from, next) {
    next(vm => vm.setData())
  },

  methods: {
    setData() {
      this.busy = true

      if (window.originURL == window.location.href && window.valid) {
        // dont make api call
        console.log('window.apiData', window.apiData)

        const { people } = window.apiData
        this.people = people
        this.busy = false
      } else {

        axios.get(window.location.href, {
          headers: {
            'AJAXRequest' : 'true',
          }
        })
        .then((res) => {
          //ajax page load
          console.log('I am ajax -- response data:', res.data)
          const { people } = res.data
          this.people = people
          this.busy = false
          window.valid = false
        })


      }
    }
  }

}
</script>

<template>
  <main>
    <TheWelcome />
    <ol>
      <li v-for="person in people">{{ person.name }}</li>
    </ol>
  </main>
</template>
