<template lang="html">
  <nav class="navbar navbar-default">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand" href="#">Todo App</a>
      </div>

      <p v-if="isLogged" class="navbar-text navbar-right">
        Hello!
      </p>
      <div v-else class="navbar-form navbar-right">
        <div class="form-group">
          <input
            type="text"
            class="form-control"
            placeholder="Username"
            v-model="userName">
        </div>
        <div class="form-group">
          <input
            type="password"
            class="form-control"
            placeholder="Password"
            v-model="pass">
        </div>
        <button type="submit" class="btn btn-default" @click="login">
          Submit
        </button>
      </div>
    </div>
</nav>
</template>

<script>
import { mapState } from 'vuex'
export default {
  computed: {
    ...mapState({
      isLogged: state => state.auth.isLogged
    })
  },
  data: function () {
    return {
      userName: '',
      pass: ''
    }
  },
  methods: {
    login: function() {
      var fd = new FormData()
      fd.append('username', this.userName)
      fd.append('password', this.pass)
      this.$store.dispatch('getToken', fd);
    }
  },
  beforeCreate: function () {
    if (window.localStorage.getItem("todo-app-storage")) {
      this.$store.dispatch('getToken')
    }
  }
}
</script>
