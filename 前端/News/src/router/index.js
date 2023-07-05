import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Sports from '../views/Sports.vue'
import Education from '../views/Education.vue'
import Politics from '../views/Politics.vue'
import Military from '../views/Military.vue'
import Science from '../views/Science.vue'
import Entertainment from '../views/Entertainment.vue'
import PersonalCenter from '../views/PersonalCenter.vue'
import Newspage from '../views/newspage.vue'
import changepassword from '../views/personal_changepassword.vue'
import MyLikes from '../views/Personal_likes.vue'
import MyComment from '../views/MyMessage_comment.vue'
import MyReply from '../views/MyMessage_reply.vue'
import Likes from '../views/MyMessage_likes.vue'
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
    path: '/politics',
    name: 'politics',
    component: Politics,
  },
  {
    path: '/military',
    name: 'military',
    component: Military,
  },
  {
    path: '/education',
    name: 'education',
    component: Education,
  },
  {
    path: '/sports',
    name: 'sports',
    component: Sports,
  },
  {
    path: '/entertainment',
    name: 'entertainment',
    component: Entertainment,
  },
  {
    path: '/science',
    name: 'science',
    component: Science,
  },
  {
    path: '/personal_center/myInfomation',
    name: 'personal_center',
    component: PersonalCenter,
  },
  {
    path: '/:channel/newspage/:title',
    name: 'newspage',
    component: Newspage,
  },
  {
    path: '/personal_center/changepassword',
    name: 'changepassword',
    component: changepassword
  },
  {
    path: '/personal_center/MyLikes',
    name: 'MyLikes',
    component: MyLikes
  },

  {
    path: '/personal_center/MyMessage/mycomment',
    name: 'MyComment',
    component: MyComment
  },
  {
    path: '/personal_center/MyMessage/myreply',
    name: 'MyReply',
    component: MyReply
  },
  {
    path: '/personal_center/MyMessage/likes',
    name: 'Likes',
    component: Likes
  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
