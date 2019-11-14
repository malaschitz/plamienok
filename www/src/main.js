import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import myMixin from './helper.js'

Vue.use(VueAxios, axios)
Vue.prototype.$axios = axios
Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,

  render: h => h(App),

  created: function () {
    this.$axios.defaults.headers.common['token'] = localStorage.token

    this.$axios.interceptors.response.use(undefined, function (error) {
      if (error.response.status === 401) {
        window.location = '/'
        return Promise.reject(error)
      }
    })

    this.$axios
      .get('/api/info')
      .then(response => {
        this.$store.version = response.data.Version
        if (response.data.User) {
          this.$store.commit('login', {
            name: response.data.User.Name,
            id: response.data.User.ID,
            email: response.data.User.Email
          })
        }
      })
      .catch(error => console.log(error))
  },

  mixins: [myMixin]
}).$mount('#app')
