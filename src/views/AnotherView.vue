<template>
    <AppHeader />
    <User :username="pageData?.user" />
    <div>
        <h1>AnotherView.vue</h1>
        
        
        <ol v-if="pageData">
            <li v-for="person in pageData.people" :key="person.name">{{ person.name }}</li>
        </ol>

        <p v-if="pageData?.message">{{ pageData.message }}</p>
    </div>
    
</template>

<script>
import User from '../components/User.vue'
import AppHeader from '../components/Header.vue'

export default {
  
  beforeRouteEnter (to, from, next) {
    next(vm => vm.pageData = window.data)
  },

  components: {
    AppHeader,
    User
  },

  data () {
    return {
      pageData: null,
    }
  },
}


</script>