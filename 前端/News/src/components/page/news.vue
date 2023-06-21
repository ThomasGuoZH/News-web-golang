<template>
    <div class="news">
        <h1>{{ article[0].title }}</h1>
        <div class="time-src">
        <p>{{ article[0].time }} {{ article[0].src }}</p>
        </div>
        <div class="picture">
            <img :src="article[0].pic" alt="image" style="text-align:center;">
        </div>
        <div class="content" v-html="article[0].content"></div>
    </div>
</template>
<script>
import { getNewsList } from '@/api/getNewsList.js';
export default {
    data() {
        return {
            // article: {
            //   title: '官方：贝林厄姆加盟皇马 转会费总价1.3亿欧元',
            //   time: '2023-06-08',
            //   src: '新浪体育讯',
            //   pic: 'https://n.sinaimg.cn/sports/transform/361/w649h512/20230608/22b2-18b5cc1aeb65892bc3319b622b0e42f7.png',
            //   content: '<p class=\"art_p\">6月8日,多特蒙德官方宣布,球队中场贝林厄姆正式转会皇马,转会费总价为1.3亿,分为1亿固定转会费,以及3000万的浮动条款。',
            // },
            article: {}
        };
    },
    computed: {
        async fetchNews() {
            return await getNewsList(1, '体育');
        },
    },
    async mounted() {
        this.article = await this.fetchNews;
        console.log(this.article);
    },
}
</script>
<style scoped>
.news {
    max-width: 800px;
    margin: 30px auto;
    padding: 20px;
}

h1 {
    font-size: 24px;
    margin-bottom: 20px;
    text-align: center;
    font-family: 'sihan-heavy',sans-serif;
}

p {
    font-size: 14px;
    margin-bottom: 2px;
    color: #888;
    text-align: center;
    font-family: 'sihan-light',sans-serif;
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
    font-family: 'sihan-regular',sans-serif;
}
</style>