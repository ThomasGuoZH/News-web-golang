<template>
    <div>
        <div v-if="user.loggedIn">
            <p>欢迎 {{ user.username }}，您已登录</p>
            <button @click="logout">退出登录</button>
            <el-link href="/personal_center/message" target="_blank">个人中心</el-link>
        </div>
        <div v-else>
            <p>您还未登录</p>
            <button @click="login">登录</button>
        </div>
    </div>
</template>

<script>
export default {
    computed: {
        user() {
            return this.$store.state.user;
        }
    },
    methods: {
        login() {
            const user = {
                username: 'test',
                avatar: 'https://example.com/avatar.png'
            };
            this.$store.dispatch('login', user);
        },
        logout() {
            this.$store.dispatch('logout');
        }
    }
};
</script>