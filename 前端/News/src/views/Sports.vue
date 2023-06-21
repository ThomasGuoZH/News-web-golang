<template>
  <div class="sports">
    <navigation :items="items">
      <template v-slot:体育>
        <b>{{ items[4].name }}</b>
      </template>
    </navigation>
    <backtotop />
    <h3 style="color:black;margin-top:50px;text-align: center; font-size: 25px;font-family: 'common-black',sans-serif;">
      体育新闻</h3>
    <div class="carousel-container">
      <el-carousel :autoplay="true" :interval="5000" height="480px">
        <el-carousel-item v-for="(news, index) in sportsNewslist.slice(0, 4)" :key="index">
          <router-link :to="'/sports/newspage/' + news.title">
            <img class="slider-image" :src="news.pic" alt="Slider Image">
            <div class="news-title">{{ news.title }}</div>
          </router-link>
          <div class="slider-caption">
            <h3>{{ news.title }}</h3>
            <p>{{ news.content }}</p>
          </div>
        </el-carousel-item>
      </el-carousel>
    </div>
    <div class="news-layout">
      <div class="view-mode">
        <el-button @click="changeViewMode('list')" :class="{ active: viewMode === 'list' }">列表浏览</el-button>
        <el-button @click="changeViewMode('card')" :class="{ active: viewMode === 'card' }">卡片浏览</el-button>
      </div>

      <div class="news-container">
        <!-- 列表模式 -->
        <div v-if="viewMode === 'list'" class="news-lists">
          <div class="news-list" v-for="(news, index) in sportsNewslist.slice(4, sportsNewslist.length)" :key="index">
            <div class="news-item">
              <div class="news-thumbnail">
                <router-link :to="'/sports/newspage/' + news.title">
                  <img class="thumbnail-image" :src="news.pic" alt="News Thumbnail">
                </router-link>
              </div>
              <div class="news-details">
                <router-link :to="'/sports/newspage/' + news.title">
                  <h4>{{ news.title }}</h4>
                </router-link>
                <div class="news-src">{{ news.src }}</div>
                <div class="news-time">{{ news.time }}</div>

              </div>
            </div>
          </div>
        </div>

        <!-- 卡片模式 -->
        <div v-else-if="viewMode === 'card'" class="news-cards">
          <div class="news-card" v-for="(news, index) in sportsNewslist.slice(4, sportsNewslist.length)" :key="index">
            <div class="card-image">
              <router-link :to="'/sports/newspage/' + news.title" target="_blank">
                <img class="card-thumbnail" :src="news.pic" alt="News Thumbnail">
              </router-link>
            </div>
            <div class="card-details">
              <router-link :to="'/sports/newspage/' + news.title" target="_blank">
                <h4>{{ news.title }}</h4>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
  
<script>
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import navigation from '../components/layout/nav.vue'
import carousel from '@/components/layout/carousel.vue';
import { getNewsList } from '@/api/getNewsList'
import { storeNews } from '@/api/news'
export default {
  name: 'sports',
  components: {
    navigation,
    sidebox,
    backtotop,
    carousel
  },
  data() {
    return {
      viewMode: 'list', // 默认为列表浏览模式
      items: [
        { name: '首页', url: '/home' },
        { name: '时政', url: '/politics' },
        { name: '科技', url: '/science' },
        { name: '娱乐', url: '/entertainment' },
        { name: '体育', url: '/sports' },
        { name: '军事', url: '/military' },
        { name: '教育', url: '/education' }],
      sportsNewslist: [],

    };
  },
  methods: {
    changeViewMode(mode) {
      this.viewMode = mode;
    }
  },
  mounted: async function () {
    this.sportsNewslist = await getNewsList(12, '体育');
    const res = await storeNews(this.sportsNewslist);
    console.log(res);
  }
};
</script>
  
<style scoped>
.sports {
  font-family: 'common-regular', sans-serif;
  background-color: #F9F7F7;
  padding-bottom: 5%;
}

.slider-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.carousel-container {
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
  margin-bottom: 20px;
}

.slider-caption {
  background-color: rgba(0, 0, 0, 0.5);
  color: #fff;
  padding: 10px;
}

a {
  text-decoration: none;
}

.news-layout {
  padding: 20px;
}

.view-mode {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin: 0px auto 10px auto;
  width: 207px;
}

.news-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  width: 1000px;
  background-color: #DBE2EF;
  margin: auto;
}

.news-lists {
  flex-basis: 100%;
}

.news-list {
  flex-basis: 100%;
  width: 900px;
  margin: auto;
}

.news-item {
  display: flex;
  width: 100%;
  padding: 10px;
}

.news-thumbnail {
  width: 30%;
  height: 144px;
  overflow: hidden;
  margin-right: 20px;
}

.thumbnail-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.news-details {
  flex: 1;
  width: 60%;
}

.news-src {
  display: inline;
  float: left;
  margin-top: 85px;
}

.news-time {
  display: inline;
  float: right;
  font-size: 15px;
  margin-top: 90px;
}

.news-details h4 {
  font-size: 20px;
  margin-bottom: 5px;
  color: black;
}


.news-cards {
  /* flex-basis: 100%; */
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  justify-content: flex-start;
  align-items: flex-start;
}

.news-card {
  flex-basis: 22.9%;
  margin-bottom: 20px;
  padding: 10px;
  margin-right: 10px;
  margin-left: 10px;
  box-sizing: border-box;
}

.card-image {
  width: 100%;
  height: 200px;
  overflow: hidden;
  margin-bottom: 10px;
}

.card-thumbnail {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.card-details {
  height: 50px;
}

.card-details h4 {
  font-size: 16px;
  color: black;
  margin: 0;
}

.empty-card {
  flex-basis: 25%;
  visibility: hidden;
  /* 隐藏占位卡片 */
}
</style>
  
  
  
  
  