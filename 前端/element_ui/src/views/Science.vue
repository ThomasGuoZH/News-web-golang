<template>
    <div class="science">
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
                    <li><a href="/science" style="font-weight: 800;">科技</a></li>
                    <li><a href="/entertainment">娱乐</a></li>
                    <li><a href="/game">游戏</a></li>
                    <li><a href="/health">健康</a></li>
                    <li><a href="/other">其他</a></li>
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
    name: 'science',
    components: {
        login_register,
        sidebox,
    }
}
</script>

<style>
@import '../assets/style/科技style.css'
</style>