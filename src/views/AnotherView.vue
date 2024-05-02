<template>
    <User :username="pageData?.user" />
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
import User from '../components/User.vue'

export default {
  
  beforeRouteEnter (to, from, next) {
    next(vm => vm.pageData = window.data)
  },

  components: {
    User
  },

  data () {
    return {
      pageData: null,
    }
  },
}


</script>