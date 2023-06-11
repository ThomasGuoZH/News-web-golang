<template>
    <div class="politics">
        <navigation :items="items">
            <template v-slot:时政>
                <b>{{ items[1].name }}</b>
            </template>
        </navigation>
        <sidebox />
        <backtotop />
        <div class="main"> 
            <el-row>
          <el-col :span="16">
            <div class="news-list">
              <div class="news-item" v-for="(item, index) in newsList" :key="index">
                <div class="news-header">
                  <img :src="item.imgUrl" alt="news-img">
                  <div class="news-title">{{ item.title }}</div>
                </div>
                <div class="news-content">{{ item.content }}</div>
                <div class="news-footer">
                  <span class="news-time">{{ item.time }}</span>
                  <el-button type="text" class="news-more">查看详情</el-button>
                </div>
              </div>
            </div>
          </el-col>
          <el-col :span="7" offset="1">
            <div class="side-bar">
              <div class="side-item">
                <div class="side-title">热门新闻</div>
                <div class="side-content">
                  <ul>
                    <li v-for="(item, index) in hotNewsList" :key="index">
                      <router-link :to="{ path: '/news/' + item.id }">{{ item.title }}</router-link>
                    </li>
                  </ul>
                </div>
              </div>
              <div class="side-item">
                <div class="side-title">推荐新闻</div>
                <div class="side-content">
                  <ul>
                    <li v-for="(item, index) in recommendNewsList" :key="index">
                      <router-link :to="{ path: '/news/' + item.id }">{{ item.title }}</router-link>
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
export default {
    name: 'politics',
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
            newsList: [
                {
                id: 1,
                title: '新闻标题1',
                content: '新闻内容1',
                imgUrl: 'https://picsum.photos/400/300?random=1',
                time: '2021-01-01'
                },
                {
                id: 2,
                title: '新闻标题2',
                content: '新闻内容2',
                imgUrl: 'https://picsum.photos/400/300?random=2',
                time: '2021-01-02'
                },
                {
                id: 3,
                title: '新闻标题3',
                content: '新闻内容3',
                imgUrl: 'https://picsum.photos/400/300?random=3',
                time: '2021-01-03'
                }],
            hotNewsList: [
                {
                id: 1,
                title: '热门新闻1'
                },
                {
                id: 2,
                title: '热门新闻2'
                },
                {
                id: 3,
                title: '热门新闻3'
                }],
            recommendNewsList: [
                {
                id: 1,
                title: '推荐新闻1'
                },
                {
                id: 2,
                title: '推荐新闻2'
                },
                {
                id: 3,
                title: '推荐新闻3'
                }],
            
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
.main{
    width:1200px;
    margin:100px auto 0;
    background-color: rgb(247, 239, 239);
}
.news-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.news-item {
  width: 48%;
  margin-bottom: 20px;
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
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.news-time {
  font-size: 12px;
  color: #999;
}

.news-more {
  font-size: 12px;
  color: #409EFF;
}

.side-bar {
  padding: 20px;
  background-color: #f5f5f5;
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