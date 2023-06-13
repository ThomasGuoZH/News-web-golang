<template>
  <div class="sports">
    <navigation :items="items">
      <template v-slot:体育>
        <b>{{ items[4].name }}</b>
      </template>
    </navigation>
    <sidebox />
    <backtotop />
    <h3 style="color:black">体育新闻</h3>
    <div class="carousel-container">
      <el-carousel :autoplay="true" :interval="5000" height="480px">
        <el-carousel-item v-for="news in sliderNews" :key="news.id">
          <a :href="news.url" target="_blank">
            <img class="slider-image" :src="news.image" alt="Slider Image">
          </a>
          <div class="slider-caption">
            <h3>{{ news.title }}</h3>
            <p>{{ news.summary }}</p>
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
          <div class="news-list" v-for="news in newsList" :key="news.id">
            <div class="news-item">
              <div class="news-thumbnail">
                <a :href="news.url" target="_blank">
                  <img class="thumbnail-image" :src="news.image" alt="News Thumbnail">
                </a>
              </div>
              <div class="news-details">
                <a :href="news.url" target="_blank">
                  <h4>{{ news.title }}</h4>
                  <p>{{ news.summary }}</p>
                </a>
              </div>
            </div>
          </div>
        </div>

        <!-- 卡片模式 -->
        <div v-else-if="viewMode === 'card'" class="news-cards">
          <div class="news-card" v-for="(news, index) in newsList" :key="news.id">
            <div class="card-image">
              <a :href="news.url" target="_blank">
                <img class="card-thumbnail" :src="news.image" alt="News Thumbnail">
              </a>
            </div>
            <div class="card-details">
              <a :href="news.url" target="_blank">
                <h4>{{ news.title }}</h4>
              </a>
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
      sliderNews: [
        {
          id: 1,
          title: 'Slider News 1',
          summary: 'Summary of slider news 1',
          image: 'https://cdn.pixabay.com/photo/2015/04/23/21/59/tree-736877_1280.jpg',
          url: 'path/to/slider-news-1-page'
        },
        {
          id: 2,
          title: 'Slider News 2',
          summary: 'Summary of slider news 2',
          image: 'https://cdn.pixabay.com/photo/2017/07/24/19/57/tiger-2535888_1280.jpg',
          url: 'path/to/slider-news-2-page'
        },
        {
          id: 3,
          title: 'Slider News 1',
          summary: 'Summary of slider news 1',
          image: 'https://cdn.pixabay.com/photo/2015/04/23/21/59/tree-736877_1280.jpg',
          url: 'path/to/slider-news-1-page'
        },
        {
          id: 4,
          title: 'Slider News 2',
          summary: 'Summary of slider news 2',
          image: 'https://cdn.pixabay.com/photo/2017/07/24/19/57/tiger-2535888_1280.jpg',
          url: 'path/to/slider-news-2-page'
        },
        // 添加更多轮播新闻项...
      ],
      newsList: [
        {
          id: 1,
          title: 'News 1',
          summary: 'Summary of news 1111111111111111111111111111111111111111111111111111111111111111111111111',
          image: 'https://cdn.pixabay.com/photo/2023/05/30/15/43/koala-8028992_1280.jpg',
          url: 'path/to/news-1-page'
        },
        {
          id: 2,
          title: 'News 2',
          summary: 'Summary of news 2',
          image: 'https://cdn.pixabay.com/photo/2020/04/23/18/41/light-5083606_1280.jpg',
          url: 'path/to/news-2-page'
        },
        {
          id: 3,
          title: 'News 3',
          summary: 'Summary of news 11111111111',
          image: 'https://cdn.pixabay.com/photo/2016/07/11/20/34/lost-places-1510592_1280.jpg',
          url: 'path/to/news-1-page'
        },
        {
          id: 4,
          title: 'News 4',
          summary: 'Summary of news 2',
          image: 'https://cdn.pixabay.com/photo/2023/06/01/12/06/snow-8033482_1280.jpg',
          url: 'path/to/news-2-page'
        },
        {
          id: 5,
          title: 'News 5',
          summary: 'Summary of news 2',
          image: 'https://cdn.pixabay.com/photo/2023/06/07/18/14/giraffes-8047856_1280.jpg',
          url: 'path/to/news-2-page'
        },
         {
          id: 6,
          title: 'News 6',
          summary: 'Summary of news 2',
          image: 'https://cdn.pixabay.com/photo/2023/05/24/14/43/cat-8015038_1280.jpg',
          url: 'path/to/news-2-page'
        },
        // 添加更多新闻项...
      ]
    };
  },
  methods: {
    changeViewMode(mode) {
      this.viewMode = mode;
    }
  }
};
</script>
  
<style>
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
  width: 900px;
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

.news-details h4 {
  font-size: 20px;
  margin-bottom: 5px;
  color: black;
}

.news-details p {
  font-size: 16px;
  color: #666;
}

.news-cards {
  /* flex-basis: 100%; */
  display: flex;
  flex-wrap: wrap;
  width: 960px;
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
  
  
  
  
  