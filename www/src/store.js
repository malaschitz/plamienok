import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({

  state: {
    name: "",
    logged: false,
    email: "",
    alert: "",
  },

  mutations: {
    login(state, data) {
      state.name = data.name;
      state.email = data.email;
      state.logged = true;
    },

    logout(state) {
      state.name = "";
      state.email = "";
      state.logged = false;
    },

    alert(state, data) {
      state.alert = data;
    }
  },

  actions: {

  }
})
