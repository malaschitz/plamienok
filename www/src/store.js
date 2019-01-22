import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({

  state: {
    name: "",
    logged: false,
    email: "",
  },

  mutations: {
    login(state, data) {
      state.name = data.name;
      state.email = data.email;
      state.logged = true;
    }
  },

  actions: {

  }
})
