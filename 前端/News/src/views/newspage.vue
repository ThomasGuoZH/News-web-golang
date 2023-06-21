<template>
  <div class="newspage">
    <navigation :items="items">
      <template v-slot:[channelname]>
        <b>{{ channelname }}</b>
      </template>
    </navigation>
    <div class="routeback">
      <el-button class= 'back' @click="routeBack" >
        <img src="../assets/icons/back.svg" alt="返回">
      </el-button>
    </div>
    <news />
    <comment />
  </div>
</template>
  
<script>
import navigation from "../components/layout/nav.vue"
import news from "../components/page/news.vue"
import comment from "../components/page/comment.vue"
export default {
  components: {
    navigation,
    news,
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
    }
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
</style>
  

  