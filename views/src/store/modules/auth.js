import * as types from '../mutation-types'
import api from '../utils/api'
import Noty from 'noty'
import Vue from 'vue'

const state = {
  isLogged: false
};

const actions = {
  getToken: function (context, data) {
    var storageToken = window.localStorage.getItem('todo-app-storage')
    if (!storageToken) {
    Vue.http.post('/api/token', data, {'Content-Type': 'application/x-www-form-urlencoded'})
      .then((res) => {
        window.localStorage.setItem('todo-app-storage', 'Bearer ' + res.body.token)
        context.commit(types.SET_LOGGED_IN)
        context.dispatch('getTodos')
      })
      .catch((error) => {
        new Noty({
          type: 'error',
          layout: 'topRight',
          text: 'Invalid credentials',
        }).show()
      })
    } else {
      context.commit(types.SET_LOGGED_IN)
      context.dispatch('getTodos')
    }
  }
};

const mutations = {
  [types.SET_LOGGED_IN]: function (state) {
    state.isLogged = true
  }
};

export default {
  state,
  actions,
  mutations
}
