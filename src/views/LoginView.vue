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
        <div v-if="errs || pageData?.errors" class="my-2 bg-red-300">
            <span :key="name" v-for="(value, name) in errs">{{ value }}</span>
            <span :key="name" v-for="(value, name) in pageData?.errors">{{ value }}</span>
        </div>

        <input type="hidden" name="csrf_token" :value="csrf_token" />

        <label>Username</label>
        <input class="bg-gray-100" v-model="auth.username" type="text" name="username" />

        <label>Password</label>
        <input class="bg-gray-100" v-model="auth.password" type="password" name="password" />
        <span v-if="errs?.password">{{ errs.password }}</span>
        <span v-if="pageData?.errors?.password">{{ pageData?.errors?.password }}</span>
        

        <SubmitButton class="my-2 bg-blue-700 text-white rounded p-1" @click="clearErrs" url="/login" :formData="auth" :onErr="setErrs">
            Login
        </SubmitButton>
        <button class="bg-cyan-300 rounded p-1" type="submit">Login non Ajax</button>
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
            console.log('setErrs', val)
            this.errs = val
        },
        clearErrs() {
            console.log('clearErrs')
            this.errs = null

            if (this.pageData?.errors) {
                this.pageData.errors = null
            }
        }
    },

    computed: {
        errors() {
            if (this.pageData?.errors) {
                return this.pageData.errors
            }

            if (this.errs) {
                return this.errs
            }

            return null
        },

        csrf_token() {
            let cookies = document.cookie.split("; ")

            for (let i=0; i<cookies.length; i++) {
                let [key, val] = cookies[i].split("=")

                if (key == "XSRF-TOKEN") {
                    return val
                }
            }
            return ""
        },
    },

    watch: {
        pageData(value) {
            if (value.user) {
                this.$router.push('/private/me')
            }
        },
    }, 
    beforeRouteEnter (to, from, next) {
        next(vm => vm.pageData = window.data)
    },

}
</script>