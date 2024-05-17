<template>
    <AppHeader />
    <main>
        <dl>
            <dt>View name</dt>
            <dd>CurrentUserView.vue</dd>
            <dt>Username</dt>
            <dd>{{ pageData?.user }}</dd>
        </dl>
    </main>
</template>
<script>
import AppHeader from '../components/Header.vue'
export default {
    components: {
        AppHeader,
    },
    data() {
        return {
            pageData: null,
        }
    },

    beforeRouteEnter (to, from, next) {
        next(vm => vm.pageData = window.data)
    }
}
</script>