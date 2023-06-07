
<template>
  <div class="home">
    <span class="top">
      <h2>新闻论坛</h2>
      <!-- vue输入用户名 -->

      <div class="user">
        <login_register />
      </div>

      <div class="nav">
        <ul>
          <li><a href="/home" style="font-weight: 800;">首页</a></li>
          <li><a href="/pe">体育</a></li>
          <li><a href="/science">科技</a></li>
          <li><a href="/entertainment">娱乐</a></li>
          <li><a href="/game">游戏</a></li>
          <li><a href="/health">健康</a></li>
          <li><a href="/other">其他</a></li>
        </ul>
      </div>
    </span>
    <h3 style="color:red">头条新闻</h3>
    <sidebox />
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
        <h4 style="color: #f97272;">体育</h4>
        <a href="体育.html" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #c4cb6d;">科技</h4>
        <a href="科技.html" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #f27fdb;">娱乐</h4>
        <a href="娱乐.html" target="_blank">更多>></a>
        <div></div>
      </li>
    </ul>
    <ul class="context">
      <li>
        <h4 style="color: #91dcfa;">游戏</h4>
        <a href="游戏.html" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #a779fb;">健康</h4>
        <a href="健康.html" target="_blank">更多>></a>
        <div></div>
      </li>
      <li>
        <h4 style="color: #9cf194;">其他</h4>
        <a href="其他.html" target="_blank">更多>></a>
        <div></div>
      </li>
    </ul>
  </div>
</template>
<script>
import login_register from '../components/login_register.vue'
import carousel from '../components/carousel.vue'
import sidebox from '../components/sidebox.vue'
import backtotop from '../components/backtotop.vue'
export default {
  name: 'home',
  components: {
    login_register,
    sidebox,
    carousel,
    backtotop,
  }
}

</script>
<style>
@import '../assets/style/首页style.css'
</style>
