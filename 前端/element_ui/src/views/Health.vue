<template>
    <div class="2">
        <span class="top">
            <h2>新闻论坛</h2>
            <!-- vue输入用户名 -->
            <div class="user">
                <login_register />
            </div>
            <div class="nav">
                <ul>
                    <li><a href="/home">首页</a></li>
                    <li><a href="/pe">体育</a></li>
                    <li><a href="/science">科技</a></li>
                    <li><a href="/entertainment">娱乐</a></li>
                    <li><a href="/game">游戏</a></li>
                    <li><a href="/health" style="font-weight: 800;">健康</a></li>
                    <li><a href="/other">其他</a></li>
                </ul>
            </div>
        </span>
        <cebianlan />
    </div>
</template>
<script>
import login_register from '../components/login_register.vue'
import cebianlan from '../components/侧边栏.vue'
export default {
    name: '2',
    components: {
        login_register,
        cebianlan
    }
}
</script>
<style>
@import '../assets/style/健康style.css'
</style>