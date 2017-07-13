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
  },
  logout: function (context) {
    window.localStorage.removeItem('todo-app-storage')
    context.commit(types.SET_LOGGED_OUT)
    context.commit(types.CLEAR_ALL_ITEMS)
  }
};

const mutations = {
  [types.SET_LOGGED_IN]: function (state) {
    state.isLogged = true
  },
  [types.SET_LOGGED_OUT]: function (state) {
    state.isLogged = false
  }
};

export default {
  state,
  actions,
  mutations
}
