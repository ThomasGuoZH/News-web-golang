<template>
    <div class="entertainment">
        <navigation :items="items">
            <template v-slot:娱乐>
                <b>{{ items[3].name }}</b>
            </template>
        </navigation>
        <backtotop />
        <h3 style="color:black;margin-top:50px;text-align: center; font-size: 25px;font-family: 'common-black',sans-serif;">娱乐新闻</h3>
        <div class="main">
            <el-row>
                <el-col :span="16">
                    <div class="news-list">
                        <div class="news-item" v-for="(news, index) in newsList.slice(6, 14)" :key="index">
                            <div class="news-header">
                                <img :src="news.pic" alt="news-img">
                                <div class="news-title">{{ news.title }}</div>
                            </div>
                            <div class="news-content" v-html="news.content"></div>
                            <div class="news-footer">
                                <span class="news-time">{{ news.time }}</span>
                                <router-link :to="'/entertainment/newspage/' + news.title" target="_blank"><el-button
                                        type="text" id="news-more">查看详情</el-button></router-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :span="6" :offset="2">
                    <div class="side-bar">
                        <div class="side-item">
                            <div class="side-title">热门新闻</div>
                            <div class="side-content">
                                <ul>
                                    <li v-for="(item, index) in newsList.slice(0, 3)" :key="index">
                                        <router-link :to="{ path: '/entertainment/newspage/' + item.id }">{{ item.title
                                        }}</router-link>
                                    </li>
                                </ul>
                            </div>
                        </div>
                        <div class="side-item">
                            <div class="side-title">推荐新闻</div>
                            <div class="side-content">
                                <ul>
                                    <li v-for="(item, index) in newsList.slice(3, 6)" :key="index">
                                        <router-link :to="{ path: '/entertainment/newspage/' + item.id }">{{ item.title
                                        }}</router-link>
                                    </li>
                                </ul>
                            </div>
                        </div>
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
    name: 'entertainment',
    components: {
        navigation,
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
            newsList: [],
        }
    },
    mounted: async function () {
        this.newsList = await getNewsList(15, '娱乐');
    }
}
</script>

<style >
.entertainment {
    font-family: 'common-regular',sans-serif;
    background-color: #F9F7F7;
    padding-bottom: 5%;
}

body {
    background-color: rgb(229, 242, 245);
}

* {
    margin: 0;
    padding: 0;
}

.main {
    width: 1200px;
    margin: 20px auto 0;
    background-color: #DBE2EF;
}

.news-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    margin-top: 5%;
    margin-left: 5%;
}

.news-item {
    width: 47%;
    margin-bottom: 20px;
    display: block;
}

.news-header {
    position: relative;
    height: 200px;
    overflow: hidden;
}

.news-header img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.news-title {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    padding: 10px;
    background-color: rgba(0, 0, 0, 0.5);
    color: #fff;
    font-size: 16px;
    line-height: 1.5;
    text-align: center;
}

.news-content {
    margin-top: 10px;
    font-size: 14px;
    line-height: 1.5;
    text-align: justify;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
}

.news-footer {
    margin-top: 10px;
    display: block;
    justify-content: space-between;
    align-items: center;
}

.news-time {
    font-size: 12px;
    color: #999;
}

#news-more {
    float: right;
    font-size: 12px;
    color: #3F72AF;
    margin-left: 300px;
}

.side-bar {
    padding: 20px;
    background-color: #e5e3e3;
}

.side-item {
    margin-bottom: 20px;
}

.side-title {
    font-size: 16px;
    font-weight: bold;
    margin-bottom: 10px;
}

.side-content ul {
    list-style: none;
    padding: 0;
    margin: 0;
}

.side-content li {
    margin-bottom: 10px;
}

.side-content a {
    color: #333;
    font-size: 14px;
    line-height: 1.5;
    text-decoration: none;
}

.side-content a:hover {
    color: #409EFF;
}
</style>