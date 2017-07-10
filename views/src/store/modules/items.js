import * as types from '../mutation-types'
import api from '../utils/api'

const state = {
  items: []
};

const actions = {
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
  deleteTodo: function (context, item) {
    var url = '/api/tasks/' + item.ID
    return api.delete(url)
      .then((response) => {
          context.commit(types.DELETE_TODO_ITEM, item.ID)
      })
      .catch((error) => {})
  }
};

const getters = {
  allItems: state => {
    return state.items;
  }
};

const mutations = {
  [types.GET_ALL_ITEMS]: function (state, items) {
    items.forEach(item => state.items.push(item))
  },
  [types.ADD_TODO_ITEM]: function (state, item) {
    state.items.push(item)
  },
  [types.DELETE_TODO_ITEM]: function (state, id) {
    var index = state.items.findIndex(item => item.ID == id)
  }
};

export default {
  state,
  getters,
  actions,
  mutations
}
