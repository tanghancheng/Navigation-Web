import Vue from 'vue'
import Router from 'vue-router'
import Navigation from '@/components/Navigation'
import Index from '@/components/Index'
import NavigationList from '@/components/Navigation/NavigationList'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/navigation',
      name: 'Navigation',
      component: Navigation
    },
    {
      path: '/navigationList',
      name: 'NavigationList',
      component: NavigationList
    },
  ]
})
