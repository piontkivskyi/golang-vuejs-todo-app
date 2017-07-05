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
    props: ['items'],
    data: function () {
      return {
        newItem: ''
      }
    },
    methods: {
      addItem: function() {
        this.$http
          .post('/api/tasks', {name: this.newItem})
          .then(
            function (res) {
              // success callback
              var body = res.body
              this.$parent.items.push(body)
            },
            function (err) {
              alert('Error while addong task occured')
            }
          ).bind(this);
        this.newItem = ''
      }
    }
  }
</script>
