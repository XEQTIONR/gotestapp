<template>
    <dl>
        <dt>User:</dt>
        <dd>{{ username ?? 'None' }}</dd>
    </dl>
</template>

<script>
export default {
    props: {
        username: String
    }
}
</script>