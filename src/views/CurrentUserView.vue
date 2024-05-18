<template>
    <AppHeader />
    <main>
        <dl>
            <dt>View name</dt>
            <dd>CurrentUserView.vue</dd>
            <dt>Username</dt>
            <dd>{{ pageData?.user }}</dd>
        </dl>
        <SubmitButton @click="clearErrs" url="/logout" :formData="null">
            Logout
        </SubmitButton>
    </main>
</template>
<script>
import AppHeader from '../components/Header.vue'
import SubmitButton from '@/components/SubmitButton.vue';
export default {
    components: {
        AppHeader,
        SubmitButton
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