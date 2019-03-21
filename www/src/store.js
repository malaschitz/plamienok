import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    name: "",
    id: "",
    logged: false,
    email: "",
    alert: ""
  },

  mutations: {
    login(state, data) {
      state.name = data.name;
      state.id = data.id;
      state.email = data.email;
      state.logged = true;
    },

    logout(state) {
      state.name = "";
      state.id = "";
      state.email = "";
      state.logged = false;
    },

    alert(state, data) {
      state.alert = data;
    }
  },

  actions: {}
});
