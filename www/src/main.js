import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from "axios";
import VueAxios from "vue-axios";

Vue.use(VueAxios, axios);
Vue.prototype.$axios = axios;
Vue.config.productionTip = false

new Vue({
  router,
  store,

  render: h => h(App),

  mounted: function() {
    this.$axios.defaults.headers.common["token"] = localStorage.token;

    this.$axios
        .get("/api/info")
        .then(response => {
          this.$store.version = response.data.Version;

          if (response.data.User) {
            this.$store.commit("login", {
              name: response.data.Student.Name,
              email: response.data.Student.Email,
            });
          }
        }).catch(error => console.log(error));
  }
}).$mount('#app')
