import { userLogin, changeInfo, changePassword } from "@/api/user";
import { Message } from 'element-ui'

const defaultCurrentUser = {
    id: '',
    username: '',
    password: '',
    sex: '',
    email: '',
    phone: '',
    token: '',
    loggedIn: false
}
export default {
    namespaced: true,
    state: {
        currentUser: defaultCurrentUser
    },
    mutations: {
        setCurrentUser(state, { userId, username, sex, email, phone, token }) {
            state.currentUser = {
                id: userId,
                username: username,
                sex: sex,
                email: email,
                phone: phone,
                token: token,
                loggedIn: true
            }
            console.log(state.currentUser);
            localStorage.setItem(`currentUser`, JSON.stringify(state.currentUser))
        },
        removeCurrentUser(state) {
            state.currentUser = defaultCurrentUser;
            localStorage.removeItem(`currentUser`);
        }
    },
    actions: {
        async login({ commit }, { username, password }) {
            const loginForm = { username, password };
            const loginResponse = await userLogin(loginForm);
            if (loginResponse.msg === '登录成功!') {
                Message.success(loginResponse.msg);
                const token = loginResponse.data.token;
                const userId = loginResponse.data.userId;
                const sex = loginResponse.data.sex;
                const email = loginResponse.data.email;
                const phone = loginResponse.data.phone;
                const username = loginResponse.data.username;
                commit('setCurrentUser', { userId, username, sex, email, phone, token });
            } else if (loginResponse.msg === '用户不存在!') {
                Message.warning(loginResponse.msg)
            } else if (loginResponse.msg === '密码错误!') {
                Message.warning(loginResponse.msg)
            }
        },
        logout({ commit }) {
            Message.success("注销成功");
            commit('removeCurrentUser');
        },
        async changeInfo({ commit }, form, token) {
            const res = await changeInfo(form, token);
            commit('setCurrentUser', {
                username: res.data.username,
                sex: res.data.sex,
                email: res.data.email,
                phone: res.data.phone,
            });
            if (res.msg == '保存成功！') {
                Message.success(res.msg);
            } else if (res.mag == '用户名已存在') {
                Message.warning(res.msg);
            }
        },
        // async changePassword({commit},form)
        async loadCurrentUser({ commit }) {
            const currentUserKey = `currentuser`;
            const currentUser = JSON.parse(localStorage.getItem(currentUserKey));
            if (currentUser) {
                commit("setCurrentUser", currentUser);
            }
        },
    }
}


