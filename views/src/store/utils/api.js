import Vue from 'vue'
import VueResource from 'vue-resource'

Vue.use(VueResource)

Vue.http.interceptors.push((request, next) => {
  if (typeof(Storage) !== "undefined") {
      var jwt = window.localStorage.getItem("todo-app-storage")
  }
  jwt = jwt ? jwt : null
  if (jwt) {
    request.headers.set('Authorization', jwt)
  }
  next()
})

export default {
  get (url, request) {
    return Vue.http.get(url, request)
      .then((response) => Promise.resolve(response))
      .catch((error) => Promise.reject(error))
  },
  post (url, request) {
    return Vue.http.post(url, request)
      .then((response) => Promise.resolve(response))
      .catch((error) => Promise.reject(error))
  },
  put (url, request) {
    return Vue.http.put(url, request)
      .then((response) => Promise.resolve(response))
      .catch((error) => Promise.reject(error))
  },
  delete (url, request) {
    return Vue.http.delete(url, request)
      .then((response) => Promise.resolve(response))
      .catch((error) => Promise.reject(error))
  }
}
