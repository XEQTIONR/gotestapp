<template>
    <div>
        <h1>AnotherView.vue</h1>
        <h2 v-if="pageData">{{ pageData.message }}</h2>
        
        <ol v-if="pageData">
            <li v-for="person in pageData.people">{{ person.name }}</li>
        </ol>
    </div>
    
</template>

<script>
import { useData } from '../composables/useData.js'

export default {
  beforeRouteEnter (to, from, next) {
    next(vm => vm.pageData = window.data)
  },

  data () {
    return {
      pageData: null,
    }
  },
}


</script>