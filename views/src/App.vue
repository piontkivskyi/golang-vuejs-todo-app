<template>
  <div class="container">
    <div class="jumbotron">
      <create-item></create-item>
      <items-container :items="items"></items-container>
    </div>
  </div>
</template>

<script>
  import CreateItem from './CreateItem.vue'
  import ItemsContainer from './ItemsContainer.vue'
  export default {
    components: {
      ItemsContainer,
      CreateItem
    },
    data: function() {
      return {
        items: []
      };
    },
  }
</script>

<style lang="scss">

</style>
