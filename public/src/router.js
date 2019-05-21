import Vue from 'vue'
import Router from 'vue-router'
// == Components / Views == //
import Home from './views/Home.vue'
import Dashboard from './views/Dashboard.vue'
import Settings from './views/Settings.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: Dashboard
    },
    {
      path: '/settings',
      name: 'settings',
      component: Settings
    }
  ]
})
