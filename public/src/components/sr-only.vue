<template>
  <div class="sr-only">
    {{text}}
  </div>
</template>

<script>
  export default {
    name: "SROnly",
    props: ['text']
  }
</script>

<style scoped>
  /* From Bootstrap */
  .sr-only {
    position: absolute;
    width: 1px;
    height: 1px;
    padding: 0;
    overflow: hidden;
    clip: rect(0, 0, 0, 0);
    white-space: nowrap;
    border: 0;
  }
</style>
