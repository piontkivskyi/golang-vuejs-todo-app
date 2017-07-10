<template>
    <div class="items-container col-md-12">
      <ul class="list-group">
        <li v-for="item in items" :key="item.ID"
          v-bind:class="{'list-group-item-success': item.isReady}"
          class="list-group-item">
          <span v-text="item.name"></span>
          <div class="edit-container right">
            <a href="#" v-if="!item.isReady" @click="toogleCompletion(item)">Complete /</a>
            <a href="#" @click="deleteItem(item.ID)">Delete</a>
          </div>
        </li>
      </ul>
    </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    computed: {
      ...mapState({
        items: state => state.items.items
      })
    },
    methods: {
      deleteItem: function (id) {
        this.$store.dispatch('deleteTodo', id)
      },
      toogleCompletion: function (item) {
        this.$store.dispatch('updateTodo', item)
      }
    },
    beforeCreate: function () {
      this.$store.dispatch('getTodos')
    }
  }
</script>

<style lang="scss">
  .right {
    float: right;
  }
</style>
