<template>
  <div class="create-item-container col-md-12">
    <div class="form-group">
      <input type="text"
        class="form-control"
        v-model="newItem"
        placeholder="Add new item"
        @keyup.enter="addItem">
    </div>
  </div>
</template>

<script>
  export default {
    data: function () {
      return {
        newItem: ''
      }
    },
    methods: {
      addItem: function() {
        this.$store.dispatch('addTodo', this.newItem)
        this.newItem = ''
      }
    }
  }
</script>
