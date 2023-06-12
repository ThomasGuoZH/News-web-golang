<!-- <template>
    <div class="sports">
        <navigation :items="items">
            <template v-slot:体育>
                <b>{{ items[4].name }}</b>
            </template>
        </navigation>
        <sidebox />
        <backtotop />
    </div>
</template>
<script>
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import navigation from '../components/layout/nav.vue'
export default {
    name: 'sports',
    components: {
        navigation,
        sidebox,
        backtotop
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
        }
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
</style> -->

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
      <div class="news-layout">
        <div class="carousel-container">
          <el-carousel :autoplay="true" :interval="5000">
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
  
        <div class="news-list">
          <div class="news-row" v-for="news in newsList" :key="news.id">
            <div class="news-item">
              <div class="news-thumbnail">
                <a :href="news.url" target="_blank">
                  <img class="thumbnail-image" :src="news.image" alt="News Thumbnail">
                </a>
              </div>
              <div class="news-details">
                <h4>{{ news.title }}</h4>
                <p>{{ news.summary }}</p>
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
          // 添加更多轮播新闻项...
        ],
        newsList: [
          {
            id: 1,
            title: 'News 1',
            summary: 'Summary of news 1',
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
          // 添加更多新闻项...
        ]
      };
    }
  };
  </script>
  
  <style>
.news-layout {
  padding: 20px;
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

.news-list {
  display: flex;
  flex-wrap: wrap;
}

.news-row {
  flex-basis: 100%;
}

.news-item {
  display: flex;
  width: calc(33.33% - 10px);
  margin-right: 10px;
  margin-left: 180px;
}

.news-thumbnail {
  width: 144px;
  height: 81px;
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
}

.news-details h4 {
  font-size: 16px;
  margin-bottom: 5px;
}

.news-details p {
  font-size: 14px;
  color: #666;
}

.slider-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
  
  
  
  
  