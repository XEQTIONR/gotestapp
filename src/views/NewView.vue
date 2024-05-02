<template>
    <div>
        <User :username="pageData?.user" />
        <span>NewView.vue</span>
    </div>
</template>

<script>
import { useData } from '../composables/useData.js'
import User from '../components/User.vue'

export default {

    components: {
        User,
    },
    
    data() {
        return {
            pageData: null,
        }
    },
    beforeRouteEnter (to, from, next) {
        next(vm => vm.pageData = window.data)
  },
}
</script>