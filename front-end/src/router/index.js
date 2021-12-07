import Vue from 'vue'
import Router from 'vue-router'
import Navigation from '@/components/Navigation/Navigation'
import Index from '@/components/Index'
import NavigationList from '@/components/Navigation/NavigationList'
import NavigationIndex from '@/components/Navigation/NavigationIndex'


Vue.use(Router)
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

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
      component: Navigation,
      children: [
        {
          path: '/navigation/index',
          name: 'navigationIndex',
          component: NavigationIndex
        },
        {
          path: '/navigation/list',
          name: 'NavigationList',
          component: NavigationList
        }
      ]
    },
  ]
})
