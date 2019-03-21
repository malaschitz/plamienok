import Vue from "vue";
import Router from "vue-router";
import Dashboard from "./views/Dashboard";
import Users from "./views/Users";
import Tasks from "./views/Tasks";
import Meds from "./views/Meds";
import Diagnoses from "./views/Diagnoses";
import Sessions from "./views/Sessions";
import Cars from "./views/Cars";
import Persons from "./views/Persons";
import Person from "./views/Person";
import Meniny from "./views/Meniny";

Vue.use(Router);

export default new Router({
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "dashboard",
      component: Dashboard
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/About.vue")
    },
    { path: "/users", name: "users", component: Users },
    { path: "/tasks", name: "tasks", component: Tasks },
    { path: "/cars", name: "cars", component: Cars },
    { path: "/meds", name: "meds", component: Meds },
    { path: "/diagnoses", name: "diagnoses", component: Diagnoses },
    { path: "/sessions", name: "sessions", component: Sessions },
    { path: "/persons", name: "persons", component: Persons },
    { path: "/person/:id", name: "person", component: Person },
    { path: "/meniny", name: "meniny", component: Meniny }
  ]
});
