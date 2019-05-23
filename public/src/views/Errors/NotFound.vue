<template>
  <h1>
    404 Not Found
  </h1>
</template>

<script>
  export default {
    name: "NotFound",
    beforeCreate () {
      if (!this.$store.state.credentials.loggedIn) {
        this.$router.push("/")
      }
    }
  }
</script>

<style scoped>

</style>
