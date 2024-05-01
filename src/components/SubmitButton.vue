<template>
    <button @click="action">
        <slot />
    </button>
</template>

<script>
    export default  {
        props: {
            formData: Object,
            url: String,
            onErr: Function,
        },
        data() {
            return {}
        },
        
        methods: {
            action(e) {
                e.preventDefault()
                axios.post(this.url, this.formData, {withCredentials: true})
                    .then(res => {
                        if (res?.data?.error) {
                            this.$router.push('/login')
                            this.onErr(res.data.error)
                        } else if (res?.data?.to) {
                            this.$router.push(res.data.to)
                        } else {
                            this.$router.push('/private/me')
                        }
                    })
                    .catch(err => {
                        console.log(err)
                    })
            }
        }
    }
</script>