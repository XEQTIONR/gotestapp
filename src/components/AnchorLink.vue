<template>
    <span class="text-red" v-if="disabled">
        <slot></slot>
    </span>
    <RouterLink v-else :to="to">
        <slot></slot>
    </RouterLink>
</template>
<script>
import { RouterLink } from 'vue-router'

export default {
    props: {
        to: {
            type: String,
            required: true
        },
        disabled: {
            type: Boolean,
            default: false
        }
    },
    components: { RouterLink }
}

</script>