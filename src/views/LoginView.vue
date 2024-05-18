<template>
<AppHeader />
<div>  
    <h3 v-if="pageData?.user">{{ pageData.user }}</h3>
    <SubmitButton v-if="pageData?.user" @click="clearErrs" url="/logout" :formData="null">
        Logout
    </SubmitButton>
    <form
        v-else
        action="/login" 
        method="POST"
    >
        <label>Username</label>
        <input v-model="auth.username" type="text" name="username" />

        <label>Password</label>
        <input v-model="auth.password" type="password" name="password" />
        <span v-if="errs?.password">{{ errs.password }}</span>
        <span v-if="pageData?.errors?.password">{{ pageData?.errors?.password }}</span>

        <SubmitButton @click="clearErrs" url="/login" :formData="auth" :onErr="setErrs">
            Login
        </SubmitButton>
    </form>
</div>
</template>

<script>

import SubmitButton from '../components/SubmitButton.vue'
import AppHeader from '../components/Header.vue'

export default {

    components: { 
        AppHeader,
        SubmitButton,
    }, 

    data () {
        return {
            pageData: null,
            errs: null,
            auth: {
                username: "",
                password: ""
            }
        }
    },

    methods: {
        setErrs(val) {
            this.errs = val
        },
        clearErrs() {
            this.errs = null

            if (this.pageData.errors) {
                this.pageData.errors = null
            }
        }
    },

    mounted() {
        this.clearErrs();
        
        if (this.pageData?.user) {
            this.$router.push('/private/me')
        }
    },

    beforeRouteEnter (to, from, next) {
        next(vm => vm.pageData = window.data)
    },
}
</script>