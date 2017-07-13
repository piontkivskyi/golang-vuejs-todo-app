import * as types from '../mutation-types'
import api from '../utils/api'
import Noty from 'noty';

const state = {
  items: []
};

const actions = {
  clearItems: function (context) {

  },
  getTodos: function (context) {
    return api.get('/api/tasks')
      .then((response) => {
        context.commit(types.GET_ALL_ITEMS, response.body)
      })
      .catch((error) => {})
  },
  addTodo: function (context, name) {
    return api.post('/api/tasks', {name: name})
      .then((response) => {
        context.commit(types.ADD_TODO_ITEM, response.body)
      })
      .catch((error) => { alert('error while adding new item') })
  },
  updateTodo: function (context, item) {
    var url = '/api/tasks/' + item.ID
    return api.post(url, {name: item.name, isReady: !item.isReady})
      .then((response) => {
        context.commit(types.UPDATE_TODO_ITEM, response.body)
      })
      .catch(error => {})
  },
  deleteTodo: function (context, id) {
    var url = '/api/tasks/' + id
    return api.delete(url)
      .then((response) => {
          context.commit(types.DELETE_TODO_ITEM, id)
      })
      .catch((error) => {
        new Noty({
            text: 'Some notification text',
        }).show();
      })
  }
};

const mutations = {
  [types.CLEAR_ALL_ITEMS]: function (state) {
    state.items = []
  },
  [types.GET_ALL_ITEMS]: function (state, items) {
    if (items) {
      // clear up
      state.items = []
      items.forEach(item => state.items.push(item))
    }
  },
  [types.ADD_TODO_ITEM]: function (state, item) {
    state.items.push(item)
  },
  [types.DELETE_TODO_ITEM]: function (state, id) {
    var index = state.items.findIndex(item => item.ID == id)
    state.items.splice(index, 1)
  },
  [types.UPDATE_TODO_ITEM]: function (state, item) {
    var index = state.items.findIndex(i => item.ID == i.ID)
    state.items.splice(index, 1, item)
  }
};

export default {
  state,
  actions,
  mutations
}
