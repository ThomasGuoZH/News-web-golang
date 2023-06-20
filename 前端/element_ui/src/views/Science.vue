<template>
    <div class="science">
        <navigation :items="items">
            <template v-slot:科技>
                <b>{{ items[2].name }}</b>
            </template>
        </navigation>
        <backtotop />
        <h3 style="color:black;margin-top:50px;text-align: center; font-size: 25px;font-family: 'common-black',sans-serif;">科技新闻</h3>
        <div class="main">
            <el-row gutter="18">
                <el-col span="16">
                    <div class="news-box">
                        <el-col span="24">
                            <div class="news-carousel">
                                <el-carousel height="500px" :interval="5000">
                                    <el-carousel-item v-for=" (news, index) in scienceNewslist.slice(0, 4)" :key="index">
                                        <div class="news-item">
                                            <router-link :to="'/science/newspage/' + news.title">
                                                <img :src="news.pic">
                                                <div class="news-title">{{ news.title }}</div>
                                            </router-link>
                                        </div>
                                    </el-carousel-item>
                                </el-carousel>
                            </div>
                        </el-col>
                        <el-col span="12">
                            <div class="m1-1">
                                <div class="news-item1">
                                    <router-link :to="'/science/newspage/' + scienceNewslist[4].title">
                                        <img :src="scienceNewslist[4].pic" alt="news-img">
                                        <div class="news-title">{{ scienceNewslist[4].title }}</div>
                                    </router-link>
                                </div>
                            </div>
                        </el-col>
                        <el-col span="12">
                            <div class="m1-1">
                                <div class="news-item1">
                                    <router-link :to="'/science/newspage/' + scienceNewslist[5].title">
                                        <img :src="scienceNewslist[5].pic" alt="news-img">
                                        <div class="news-title">{{ scienceNewslist[5].title }}</div>
                                    </router-link>
                                </div>
                            </div>
                        </el-col>
                    </div>
                </el-col>
                <el-col span="8">
                    <div class="side">
                        <el-col span="24">
                            <div class="m2">
                                <div class="news-item2">
                                    <ul>
                                        <li v-for="(news, index) in scienceNewslist.slice(6, 12)" :key="index">
                                            <el-row>
                                                <router-link :to="'/science/newspage/' + news.title">
                                                    <el-col span="13">
                                                        <img :src="news.pic" alt="news-img">
                                                    </el-col>
                                                    <el-col span="11">
                                                        <div class="news-title1">{{ news.title }}</div>
                                                    </el-col>
                                                </router-link>
                                                <div class="news-time">{{ news.time }}</div>
                                            </el-row>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </el-col>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<script>
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import navigation from '../components/layout/nav.vue'
import { getNewsList } from '@/api/getNewsList'
export default {
    name: 'science',
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
            scienceNewslist: []
        }
    },
    mounted: async function () {
        this.scienceNewslist = await getNewsList(12, '科技');
    }
}

</script>

<style scoped>
.science {
    font-family: 'common-regular',sans-serif;
    background-color: #F9F7F7;
    padding-bottom: 5%;
}

.el-col {
    margin-bottom: 20px;
    border-radius: 4px;
}

.el-col:last-child {
    margin-bottom: 0;
}

.main {
    width: 1200px;
    margin: 20px auto 0;
    background-color: #DBE2EF;
    padding-top: 10px;
    padding-right: 5px;
}

.m1-1 {
    height: 200px;
    background-color: #DBE2EF;
}

.m2 {
    height: 820px;
    background-color:#3F72AF;
}

.news-carousel {
    position: relative;
    height: 600px;
    width: 770px;
    display: block;
    margin: auto;
}

.news-item {
    position: relative;
    overflow: hidden;
    height: 500px;
    width: 769.99px;
}

.news-item img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.news-item1 {
    position: relative;
    overflow: hidden;
    height: 200px;
    width: 376.68px;
}

.news-item1 img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.news-title {
    position: absolute;
    bottom: 0px;
    left: 0;
    width: 100%;
    padding: 10px;
    background-color: rgba(0, 0, 0, 0.5);
    color: #fff;
    font-size: 16px;
    line-height: 1.5;
    text-align: center;
}

.el-carousel {
    width: 770px;
    height: 550px;
    background-color: rgb(226, 224, 224);
}

.el-carousel__item {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 500px;
    width: 770px;
    margin: 0;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
}

.news-item2 ul {
    padding: 5px;;
}

.news-item2 li {
    list-style: none;
    position: relative;
    overflow: hidden;
    width: 366px;
    display: block;
    margin-bottom: 10px;
    margin-right: 5px;
}

.news-item2 div {
    margin: auto;
}

.news-item2 img {
    width: 190px;
    height: 120px;
}

.news-title1 {
    color: black
}

.news-time {
    color: rgb(61, 61, 61);
    position: absolute;
    bottom: 5px;
    left: 200px;
}
</style>