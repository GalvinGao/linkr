import Vue from 'vue'
import Router from 'vue-router'
// == Components / Views == //
import Home from './views/Home'
import Links from './views/Links'
import Settings from './views/Settings'
import NotFound from './views/Errors/NotFound'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/links',
      name: 'Links',
      component: Links
    },
    {
      path: '/settings',
      name: 'Settings',
      component: Settings
    },
    {
      path: '*',
      name: '404',
      component: NotFound
    }
  ]
})
