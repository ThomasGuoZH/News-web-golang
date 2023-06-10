<template>
  <div class="newspage">
    <navigation :items="items">
      <template v-slot:[channelname]>
        <b>{{ channelname }}</b>
      </template>
    </navigation>
    <div class="routeback">
      <el-page-header @back="routeBack" content="详情页面">
      </el-page-header>
    </div>

    <!-- <news /> -->
  </div>
</template>
  
<script>
import navigation from "../components/layout/nav.vue"
import news from "../components/page/news.vue"
export default {
  components: {
    navigation,
    news,
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
        console.log(1);
        return '时政';
      } else if (this.channel === 'sports') {
        return 4;
      } else if (this.channel === 'military') {
        return 5;
      } else if (this.channel === 'entertainment') {
        return 3;
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
    }
  },
};
</script>
<style>
.routeback {
  margin-top: -30px;
  color: white;
}
</style>
  

  