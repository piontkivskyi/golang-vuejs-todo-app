import Vue from 'vue'
import App from './components/App.vue'
import store from './store'
import VueResourse from 'vue-resource'

Vue.use(VueResourse)

new Vue({
  el: '#app',
  store,
  render: h => h(App)
})
