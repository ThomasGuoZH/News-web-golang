import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Pe from '../views/Sports.vue'
import Game from '../views/Game.vue'
import Other from '../views/Other.vue'
import Health from '../views/Health.vue'
import Science from '../views/Science.vue'
import Entertainment from '../views/Entertainment.vue'
import personal_center from '../views/personal_center.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/sports',
    name: 'sports',
    component: Pe
  },
  {
    path: '/health',
    name: 'health',
    component: Health
  },
  {
    path: '/other',
    name: 'other',
    component: Other
  },
  {
    path: '/game',
    name: 'game',
    component: Game
  },
  {
    path: '/entertainment',
    name: 'entertainment',
    component: Entertainment
  },
  {
    path: '/science',
    name: 'science',
    component: Science
  },
  {
    path:'/personal_center/myMessage',
    name: 'personal_center',
    component: personal_center
  },
  {

  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
