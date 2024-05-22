<template>
<AppHeader />
<div>  
    <h3 v-if="pageData?.user">{{ pageData.user }}</h3>
    <SubmitButton v-if="pageData?.user" @click="clearErrs" url="/logout" :formData="null">
        Logout
    </SubmitButton>
    <form
        class="flex flex-col"
        v-else
        action="/login" 
        method="POST"
    >
        <span v-if="errs?.csrf">{{ errs.csrf }}</span>
        <label>Username</label>
        <input class="bg-gray-100" v-model="auth.username" type="text" name="username" />

        <label>Password</label>
        <input class="bg-gray-100" v-model="auth.password" type="password" name="password" />
        <span v-if="errs?.password">{{ errs.password }}</span>
        <span v-if="pageData?.errors?.password">{{ pageData?.errors?.password }}</span>

        <SubmitButton class="mt-2 bg-blue-700 text-white rounded p-1" @click="clearErrs" url="/login" :formData="auth" :onErr="setErrs">
            Login
        </SubmitButton>
        <button type="submit">Login non Ajax</button>
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