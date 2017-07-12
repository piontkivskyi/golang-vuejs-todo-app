<template>
  <div class="container">
    <login></login>
    <create-item></create-item>
    <items-container></items-container>
  </div>
</template>

<script>
  import CreateItem from './CreateItem.vue'
  import ItemsContainer from './ItemsContainer.vue'
  import Login from './Login.vue'
  export default {
    components: {ItemsContainer, CreateItem, Login},
  }
</script>

<style lang="scss">

</style>
