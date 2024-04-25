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
        <input type="text" name="username" />

        <label>Password</label>
        <input type="password" name="password" />

        <button type="submit">Submit</button>

    </form>
</div>
</template>

<script>
import { useData } from '../composables/useData.js'

export default {
    data () {
        return {
            busy: false,
            pageData: null,
        }
    },

    beforeRouteEnter (to, from, next) {
        next(vm => {
            vm.setData()
        })
    },

    methods: {
        setData() {
        this.busy = true
        useData()
            .then(response => {
            this.pageData = response
            this.busy = false
            })
        }
    }
}
</script>