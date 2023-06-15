import Vue from 'vue'
import Vuex from 'vuex'
import user from './module/user.js'
import like from './module/like.js'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    user,
    like
  }
})



