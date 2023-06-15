<template>
    <span class="top">
        <h2>新闻论坛</h2>
        <div class="user">
            <User />
            <!-- <is_login /> 还未修改，后面可以把登陆注册的组件一起合并了 -->
            <!-- <el-link href="/personal_center/myMessage" target="_blank">个人中心</el-link> -->
        </div>
        <div class="nav">
            <ul>
                <li v-for="(item, index) in items" :key="index">
                    <router-link :to="item.url">
                        <template v-if="$slots[item.name]">
                            <slot :name="item.name"></slot>
                        </template>
                        <template v-else>
                            {{ item.name }}
                        </template>
                    </router-link>
                </li>
            </ul>
        </div>
    </span>
</template>
<script>
import User from '../user/user.vue'
export default {
    components: {
        User,
    },
    props: {
        items: {
            type: Array,
            required: true
        }
    }
}
</script>
<style>
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
    margin-top: -7px;
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

.nav li a {
    text-decoration: none;
    color: rgb(0, 0, 0);
    font-size: 20px;
}

.nav a:hover {
    color: blue;
}
</style>