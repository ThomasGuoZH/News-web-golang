import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    user: {
      loggedIn: false,
      username: '',
      avatar: '',
      comment: [],
      Token: ''
    }
  },
  mutations: {
    setUser(state, user) {
      state.user = user;
    },
    addComment(user, comment) {
      user.comments.push(comment);
    },
    login(state, user) {
      state.user.loggedIn = true;
      state.user.username = user.username;
      state.user.avatar = user.avatar;
    },
    logout(state) {
      state.user.loggedIn = false;
      state.user.username = '';
      state.user.avatar = '';
    }
  },
  actions: {
    login({ commit }, user) {
      // 模拟登录请求
      setTimeout(() => {
        commit('login', user);
      }, 1000);
    },
    logout({ commit }) {
      // 模拟退出登录请求
      setTimeout(() => {
        commit('logout');
      }, 1000);
    }
  }
});

