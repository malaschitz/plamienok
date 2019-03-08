import Vue from 'vue';
import Router from 'vue-router';
import Dashboard from "./views/Dashboard";
import Users from "./views/Users";
import Meds from "./views/Meds";
import Diagnoses from "./views/Diagnoses";
import Cars from "./views/Cars";
import Persons from "./views/Persons"
import Person from "./views/Person"

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {path: '/users', name:'users', component: Users},
    {path: '/cars', name:'cars', component: Cars},
    {path: '/meds', name:'meds', component: Meds},
    {path: '/diagnoses', name:'diagnoses', component: Diagnoses},
    {path: '/persons', name: 'persons', component: Persons},
    {path: '/person/:id', name: 'person', component: Person}

  ]
})
