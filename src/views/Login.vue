<template>
    <div>
    <form 
        v-if="pageData?.user"
        action="/logout"
        method="POST"
    >
        <h3>{{ pageData.user }}</h3>
        <button type="submit">Logout</button>
    </form>
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

        <SubmitButton @click="clearErrs" url="/login" :formData="auth" :onErr="setErrs" />
    </form>
</div>
</template>

<script>

import SubmitButton from '../components/SubmitButton.vue'

export default {

    components: { SubmitButton }, 

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
    },

    beforeRouteEnter (to, from, next) {
        next(vm => vm.pageData = window.data)
    },
}
</script>