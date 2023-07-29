import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/TopPage.vue'
import PlayBoard from './components/PlayBoard.vue'
import SettingServer from './components/SettingServer.vue'

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
      path: '/playBoard',
      name: 'PlayBoard',
      component: PlayBoard
    },
    {
      path: '/setting',
      name: 'SettingServer',
      component: SettingServer
    },
  ]
})