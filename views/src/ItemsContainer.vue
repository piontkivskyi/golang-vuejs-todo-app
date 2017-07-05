<template>
    <div class="items-container col-md-12">
      <ul class="list-group">
        <li v-for="item in items"
          :key="item.ID"
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
  export default {
    props: ['items'],
    methods: {
      deleteItem: function (id) {
        var url = "/api/tasks/" + id
        this.$http.delete(url).then(
          (res) => {
            this.$parent.items.forEach((item, i) => {
              if (item.ID == id) {
                this.$parent.items.splice(i, 1)
              }
            })
          },
          (err) => {
            alert('Error while deleting item')
          }
        ).bind(this)
      },
      toogleCompletion: function (item) {
        var id = item.ID
        var url = "/api/tasks/" + id
        this.$http.post(url, {name: item.name, isReady: !item.isReady})
          .then(
            (res) => {
              this.$parent.items.forEach((item, i) => {
                if (item.ID == id) {
                  this.$parent.items.splice(i, 1, res.body)
                }
              })
            },
            (err) => {
              alert('Error')
            }
          )
      }
    },
    beforeCreate: function () {
      this.$http.get('/api/tasks').then(function(response) {
        response.body.forEach((item) => {
          this.$parent.items.push(item)
        })
      }).bind(this)
    }
  }
</script>

<style lang="scss">
  .right {
    float: right;
  }
</style>
