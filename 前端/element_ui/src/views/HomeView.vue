<template>
  <div class="home">
    <navigation :items="items">
      <template v-slot:首页>
        <b>{{ items[0].name }}</b>
      </template>
    </navigation>
    <h3 style="color:brown">头条新闻</h3>
    <carousel />
    <sidebox />
    <h3>热点新闻</h3>
    <ul class="newslist">
      <li>
        <h4 style="color: black;">时政</h4>
        <router-link to="/politics" target="_blank" class="more">更多>></router-link>
        <div>
          <ul>
            <li v-for="(news, index) in politicsNewslist" :key="index" class="newstitle">
              <router-link :to="'/politics/newspage/' + news.title" target="_blank">{{ news.title }}</router-link>
            </li>
          </ul>
        </div>
      </li>
      <li>
        <h4 style="color: black;">科技</h4>
        <router-link to="/science" target="_blank" class="more">更多>></router-link>
        <div>
          <ul>
            <li v-for="(news, index) in scienceNewslist" :key="index" class="newstitle">
              <router-link :to="'/science/newspage/' + news.title" target="_blank">{{ news.title }}</router-link>
            </li>
          </ul>
        </div>
      </li>
      <li>
        <h4 style="color: black;">娱乐</h4>
        <router-link to="/entertainment" target="_blank" class="more">更多>></router-link>
        <div>
          <ul>
            <li v-for="(news, index) in educationNewslist" :key="index" class="newstitle">
              <router-link :to="'/entertainment/newspage/' + news.title" target="_blank">{{ news.title }}</router-link>
            </li>
          </ul>
        </div>
      </li>
    </ul>
    <ul class="newslist">
      <li>
        <h4 style="color: black;">体育</h4>
        <router-link to="/sports" target="_blank" class="more">更多>></router-link>
        <div>
          <ul>
            <li v-for="(news, index) in sportsNewslist" :key="index" class="newstitle">
              <router-link :to="'/sports/newspage/' + news.title" target="_blank">{{ news.title }}</router-link>
            </li>
          </ul>
        </div>
      </li>
      <li>
        <h4 style="color: black;">军事</h4>
        <router-link to="/military" target="_blank" class="more">更多>></router-link>
        <div>
          <ul>
            <li v-for="(news, index) in militaryNewslist" :key="index" class="newstitle">
              <router-link :to="'/military/newspage/' + news.title" target="_blank">{{ news.title }}</router-link>
            </li>
          </ul>
        </div>
      </li>
      <li>
        <h4 style="color: black;">教育</h4>
        <router-link to="/education" target="_blank" class="more">更多>></router-link>
        <div>
          <ul>
            <li v-for="(news, index) in educationNewslist" :key="index" class="newstitle">
              <router-link :to="'/education/newspage/' + news.title" target="_blank">{{ news.title }}</router-link>
            </li>
          </ul>
        </div>
      </li>
    </ul>
    <backtotop />
    <backtotop />
  </div>
</template>
<script>
import carousel from '../components/layout/carousel.vue'
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import navigation from '../components/layout/nav.vue'
import { getNewsList } from '@/api/getNewsList'
export default {
  data() {
    return {
      name: 'home',
      backtotopname: 'backtotop'
    }
  },
  components: {
    navigation,
    carousel,
    sidebox,
    backtotop,
  },
  data() {
    return {
      items: [
        { name: '首页', url: '/home' },
        { name: '时政', url: '/politics' },
        { name: '科技', url: '/science' },
        { name: '娱乐', url: '/entertainment' },
        { name: '体育', url: '/sports' },
        { name: '军事', url: '/military' },
        { name: '教育', url: '/education' }],
      politicsNewslist: {},
      scienceNewslist: {},
      entertainmentNewslist: {},
      sportsNewslist: {},
      militaryNewslist: {},
      educationNewslist: {}
    };
  },
  mounted: async function () {
    this.politicsNewslist = await getNewsList(8, '政治');
    this.scienceNewslist = await getNewsList(8, '科技');
    this.sportsNewslist = await getNewsList(8, '体育');
    this.militaryNewslist = await getNewsList(8, '军事');
    this.educationNewslist = await getNewsList(8, '教育');
    this.entertainmentNewslist = await getNewsList(8, '娱乐');
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
  color: brown;
}

.newslist {
  width: 1300px;
  height: 600px;
  margin: auto;
  list-style: none;
  font-family: 'common-black',sans-serif;
}

.newslist li {
  float: left;
}

.newslist li div {
  width: 400px;
  height: 550px;
  margin: 20px 0px 5px 20px;
  margin-top: 5px;
  text-align: center;
  background-color: rgb(253, 238, 215);
  border-radius: 5%;
}

.newslist li h4 {
  display: block;
  width: 400px;
  height: 30px;
  margin-top: 25px;
  border-radius: 5%;
  text-align: center;
  text-decoration: none;
  font-size: 20px;
}

.more {
  display: block;
  width: 380px;
  margin: -20px 0px 0px 30px;
  padding-bottom: 2px;
  border-bottom: 3px solid #785bf5;
  text-align: right;
  text-decoration: none;
}

.newslist li ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.newstitle {
  width: 400px;
  height: 62px;
  margin-top: 5px;
  text-align: left;
  display: flex;
  align-items: center;
  justify-content: center;
}

.newstitle a {
  font-size: 14px;
  display: block;
  margin: 0px 40px;
  width: 80%;
  text-decoration: none;
  color: #333;
  margin: auto 2px;
}

.newstitle:hover {
  color: #ff6937;
  background-color: rgb(209, 248, 234);
}
</style>
