<template>
  <div class="newspage">
    <navigation :items="items">
      <template v-slot:item-首页>
        {{ items[0].name }}
      </template>
    </navigation>

    <news />
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
    }
  }
};
</script>
  

  