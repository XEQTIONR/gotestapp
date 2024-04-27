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

export default {
    data () {
        return {
            pageData: null,
        }
    },

    beforeRouteEnter (to, from, next) {
        next(vm => vm.pageData = window.data)
    },
}
</script>