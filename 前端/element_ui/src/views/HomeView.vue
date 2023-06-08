
<template>
  <div class="home">
    <span class="top">
      <h2>新闻论坛</h2>
      <!-- vue输入用户名 -->

      <div class="user">
        <login_register />
        <!-- <is_login /> 还未修改，后面可以把登陆注册的组件一起合并了 -->
        <!-- <el-link href="/personal_center/myMessage" target="_blank">个人中心</el-link> -->
      </div>

      <div class="nav">
        <ul>
          <li><a href="/home" style="font-weight: 800;">首页</a></li>
          <li><a href="/politics">时政</a></li>
          <li><a href="/science">科技</a></li>
          <li><a href="/entertainment">娱乐</a></li>
          <li><a href="/sports">体育</a></li>
          <li><a href="/military">军事</a></li>
          <li><a href="/education">教育</a></li>
        </ul>
      </div>
    </span>
    <h3 style="color:red">头条新闻</h3>
    <carousel />
  <!-- <div class="middle" id="photo" value="1">
      <h4>今日头条</h4>
      <img src="../assets/image/1.png" alt="" id="img">
      <p class="头条字体"><a href="https://fans.sports.qq.com/post.htm?id=1766742676465516666&mid=69#1_allWithElite"
          class="review" id="t1" target="_blank">东决:绿军大胜热火总分1-3 塔图姆33+11+7破纪录巴特勒29分</a></p>
      <span class="radio">
        <input type="radio" name="select" onclick="changeview(1)" checked="checked">
        <input type="radio" name="select" onclick="changeview(2)">
        <input type="radio" name="select" onclick="changeview(3)">
        <input type="radio" name="select" onclick="changeview(4)">
        <input type="radio" name="select" onclick="changeview(5)">
      </span>
      <a href="#" class="prev" onclick="changeprev()"> &lt; </a>
      <a href="#" class="next" onclick="changenext()"> &gt; </a>
                                                                                          </div> -->

    <!-- vue输入文章 -->
    <sidebox />

    <h3>热点新闻</h3>
    <ul class="context">
      <li>
        <h4 style="color: #f97272;">时政</h4>
        <a href="/politics" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #c4cb6d;">科技</h4>
        <a href="/science" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #f27fdb;">娱乐</h4>
        <a href="/entertainment" target="_blank">更多>></a>
        <div></div>
      </li>
    </ul>
    <ul class="context">
      <li>
        <h4 style="color: #91dcfa;">体育</h4>
        <a href="/sports" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #a779fb;">军事</h4>
        <a href="/military" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #9cf194;">教育</h4>
        <a href="/education" target="_blank">更多>></a>
        <div></div>
      </li>
    </ul>
    <backtotop></backtotop>
  </div>
</template>
<script>
import login_register from '../components/user/login_register.vue'
import carousel from '../components/layout/carousel.vue'
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import is_login from '@/components/user/is_login.vue'
export default {
  data() {
    return {
      name: 'home',
      backtotopname: 'backtotop'
    }
  },
  components: {
    login_register,
    carousel,
    sidebox,
    backtotop,
    is_login
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

.nav a {
  text-decoration: none;
  color: rgb(0, 0, 0);
  font-size: 20px;
}

.nav a:hover {
  color: blue;
}



.middle {
  width: 500px;
  height: 400px;
  margin: 150px auto;
  position: relative;
}

.middle h4 {
  text-align: center;
  font-size: 25px;
  margin: -80px 0px 20px 0px;
  color: red;
  /* 加样式 */
}

.middle img {
  width: 100%;
}

.头条字体 {
  text-align: center;
  display: block;

}

.头条字体 a {
  display: block;
  text-decoration: none;
  text-align: center;
  width: 100%;
  height: 20px;
  color: black;
}

.头条字体 a:hover {
  color: blue;
  background-color: azure;
}

.radio {
  position: absolute;
  top: 90%;
  left: 40%;
}


h3 {
  margin: 50px 0px 25px 0px;
  text-align: center;
  font-size: 25px;
  color: blueviolet;
}

.context {
  width: 1300px;
  height: 600px;
  margin: 20px auto;
  list-style: none;
}

.context li {
  float: left;
}

.context li div {
  width: 400px;
  height: 550px;
  background-color: white;
  margin: 20px 0px 5px 20px;
  margin-top: 5px;
  text-align: center;
  border-radius: 5%;
}

.context li h4 {
  display: block;
  width: 400px;
  height: 30px;
  margin-top: 25px;
  border-radius: 5%;
  text-align: center;
  text-decoration: none;
  font-size: 20px;
}

.context li a {
  display: block;
  width: 380px;
  margin: -20px 0px 0px 30px;
  padding-bottom: 2px;
  border-bottom: 3px solid #785bf5;
  text-align: right;
  text-decoration: none;
}
</style>
