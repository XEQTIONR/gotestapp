<template>
    <header class="w-full p-3 flex justify-start items-center">
      <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="75" height="75" />
        <h1 class="text-xl">Go Vue App</h1>
        <nav>
          <AnchorLink class="ml-6 font-bold" to="/">Home</AnchorLink>
          <AnchorLink class="ml-6 font-bold" to="/about">About</AnchorLink>
          <AnchorLink class="ml-6 font-bold" to="/private/another">Another</AnchorLink>
          <AnchorLink class="ml-6 font-bold" to="/login">Login</AnchorLink>
          <AnchorLink class="ml-6 font-bold" to="/register">Register</AnchorLink>
          <AnchorLink class="ml-6 font-bold" to="/private/new">New</AnchorLink>
          <AnchorLink class="ml-6 font-bold" to="/private/me">Me</AnchorLink>
        </nav>
    </header>
</template>

<script>
import AnchorLink from './AnchorLink.vue';

export default {
    components: { AnchorLink },
}
</script>