import Vue from 'vue'
import Router from 'vue-router'
import Navigation from '@/components/Navigation/Navigation'
import Index from '@/components/Index'
import NavigationList from '@/components/Navigation/NavigationList'
import NavigationIndex from '@/components/Navigation/NavigationIndex'
import NoteIndex from '@/components/StudyNote/NoteIndex'
import NoteEditor from '@/components/StudyNote/NoteEditor'
import NoteHistoryList from '@/components/StudyNote/NoteHistoryList'
import NoteDetail from '@/components/StudyNote/NoteDetail'
import NoteEditList from '@/components/StudyNote/NoteEditList'


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
          path: '/',
          name: 'navigationIndex',
          component: NavigationIndex
        },
        {
          path: '/navigation/list',
          name: 'NavigationList',
          component: NavigationList
        },
      ]
    },
    {
      path: "/note",
      name: 'note',
      component: NoteIndex,
      children: [
        {
          path: "/",
          name: "noteHistoryList",
          component: NoteHistoryList,
        },
        {
          path: "/noteEditor/:id",
          name: "noteEditor",
          component: NoteEditor,
        }, 
        {
          path: "/noteEditList",
          name: "NoteEditList",
          component: NoteEditList,
        }, 
        {
          path: "/noteDetail/:id",
          name: "noteDetail",
          component: NoteDetail,
        },

      ]
    }

  ]
})
