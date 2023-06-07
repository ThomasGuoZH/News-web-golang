<template>
    <div class="3">
        <span class="top">
            <h2>新闻论坛</h2>
            <!-- vue输入用户名 -->
            <div class="user">
                <login_register />
            </div>
            <div class="nav">
                <ul>
                    <li><a href="/home">首页</a></li>
                    <li><a href="/sports">体育</a></li>
                    <li><a href="/science">科技</a></li>
                    <li><a href="/entertainment">娱乐</a></li>
                    <li><a href="/game">游戏</a></li>
                    <li><a href="/health">健康</a></li>
                    <li><a href="/other" style="font-weight: 800;">其他</a></li>
                </ul>
            </div>
        </span>
        <sidebox />
    </div>
</template>
<script>
import login_register from '../components/login_register.vue'
import sidebox from '../components/sidebox.vue'
export default {
    name: '3',
    components: {
        login_register,
        sidebox,
    }
}
</script>
<style>
body {
    background-color: rgb(229, 242, 245);
}

* {
    margin: 0;
    padding: 0;
}

.top {
    background-color: #333333;
    color: #ffffff;
    /* 调背景样式 */
    display: block;
    height: 108px;
    border-collapse: collapse;
}

h2 {
    color: rgb(255, 255, 255);
    padding: 15px 0 0 30px;
    text-align: left;
    font-size: 30px;
    /* 加样式 */
}

.user {
    color: #ffffff;
    text-align: right;
    margin-right: 20px;
}

.nav {
    text-decoration: none;
    background-color: white;
}

.nav ul {
    display: flex;
    list-style: none;
    margin-top: 20px;
    border-bottom: black solid 3px;
}

.nav li {
    margin: auto;
}

.nav a {
    text-decoration: none;
    color: rgb(0, 0, 0);
    font-size: 20px;
}

.nav a:hover {
    color: blue;
}

.fixed {
    position: fixed;
    left: 50%;
    margin-left: 650px;
    top: 200px;
    width: 100px;
    height: 200px;
    background-color: #b9f1f1;
    text-align: center;
}
</style>