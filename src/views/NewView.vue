<template>
    NewView.vue
</template>

<script>
import { useData } from '../composables/useData.js'

export default {
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