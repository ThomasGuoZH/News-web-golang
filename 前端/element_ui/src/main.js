import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import axios from 'axios'
Vue.prototype.$axios = axios
import VueAxios from 'vue-axios';
import './global.css';

Vue.use(VueAxios, axios)
Vue.config.productionTip = false
// 配置请求的根路径
axios.defaults.baseURL = '/api'
// 挂载到原型对象之前 先设置拦截器 通过axios请求拦截器添加token，保证拥有获取数据的权限
// axios.interceptors.request.use(config => {
//   //在 request 拦截器中， 展示进度条 NProgress.start() 
//   NProgress.start()
//   // 为请求头对象添加Token验证的Authorization字段
//   config.headers.Authorization = window.sessionStorage.getItem('token')
//   // 最后都必须 return config
//   return config
// })
// // 在 response 拦截器中， 隐藏进度条 NProgress.done()
// axios.interceptors.response.use(config => {
//   NProgress.done()
//   return config
// })

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
