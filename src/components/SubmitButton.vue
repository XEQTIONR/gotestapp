<template>
    <button @click="action">
        <slot />
    </button>
</template>

<script>
import axios from 'axios'

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
                    if (res?.data?.to) {
                        this.$router.push(res.data.to)
                    } else {
                        this.$router.push('/')
                    }
                })
                .catch(err => {
                    if (this.onErr) {
                        this.onErr(err?.response?.data?.error)
                    }
                })
        }
    }
}
</script>