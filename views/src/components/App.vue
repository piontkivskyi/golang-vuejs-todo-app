<template>
  <div class="container">
    <create-item></create-item>
    <items-container></items-container>
  </div>
</template>

<script>
  import CreateItem from './CreateItem.vue'
  import ItemsContainer from './ItemsContainer.vue'
  export default {
    components: {ItemsContainer, CreateItem},
  }
</script>

<style lang="scss">

</style>
