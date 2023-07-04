<template>
  <div class="newspage">
    <navigation :items="items">
      <template v-slot:[channelname]>
        <b>{{ channelname }}</b>
      </template>
    </navigation>
    <div class="routeback">
      <el-button class='back' @click="routeBack">
        <img src="../assets/icons/back.svg" alt="返回">
      </el-button>
    </div>
    <div class="news">
      <h1>{{ news.title }}</h1>
      <div class="time-src">
        <p>{{ news.time }} {{ news.src }}</p>
      </div>
      <div class="picture">
        <img :src="news.pic" alt="image" style="text-align:center;">
      </div>
      <div class="content" v-html="news.content"></div>
    </div>
    <comment />
  </div>
</template>
  
<script>
import navigation from "../components/layout/nav.vue"
import comment from "../components/page/comment.vue"
import { getNews } from '@/api/news';
export default {
  components: {
    navigation,
    comment,
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
      channel: '',
      pathFromUrl: '',
      news: {
        title: '',
        content: '',
        src: '',
        time: '',
        pic: '',
      }
    }
  },
  computed: {
    getchannel() {
      if (this.channel === '' || this.channel === 'home') {
        return '首页';
      } else if (this.channel === 'education') {
        return '教育';
      } else if (this.channel === 'science') {
        return '科技';
      } else if (this.channel === 'politics') {
        return '时政';
      } else if (this.channel === 'sports') {
        return '体育';
      } else if (this.channel === 'military') {
        return '军事';
      } else if (this.channel === 'entertainment') {
        return '娱乐';
      }
    },
    channelname() {
      return this.getchannel;
    },
    async fetchNews() {
      this.news.title = this.$route.params.title;
      return await getNews(this.news.title);
    },
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.pathFromUrl = from.fullPath;
      vm.channel = to.params.channel;
    })
  },
  methods: {
    routeBack() {
      this.$router.push({ path: this.pathFromUrl })
      window.close();
    }
  },
  async mounted() {
    console.log(this.fetchNews);
    this.news = await this.fetchNews;
    console.log(this.news);
  },
};
</script>
<style scoped>
.routeback {
  margin-top: 30px;
  margin-left: 15px;
}

.back {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.back img {
  height: 25px;
  width: 25px;
}

.news {
  max-width: 800px;
  margin: 30px auto;
  padding: 20px;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
  font-family: 'sihan-heavy', sans-serif;
}

p {
  font-size: 14px;
  margin-bottom: 2px;
  color: #888;
  text-align: center;
  font-family: 'sihan-light', sans-serif;
}

img {
  width: 60%;
  display: block;
  margin: 20px auto 30px auto;
}

.content {
  font-size: 16px;
  line-height: 1.6;
  text-indent: 2em;
  font-weight: 500px;
  font-family: 'sihan-regular', sans-serif;
}
</style>
  

  