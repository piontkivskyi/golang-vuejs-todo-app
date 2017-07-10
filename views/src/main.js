import Vue from 'vue'
import Vuex from 'vuex'
import App from './components/App.vue'
import store from './store'
import VueResourse from 'vue-resource'

Vue.use(VueResourse)
Vue.use(Vuex)

new Vue({
  el: '#app',
  store,
  render: h => h(App)
})
