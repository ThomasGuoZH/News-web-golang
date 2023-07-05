<template>
  <div class="newspage">
    <navigation :items="items">
      <template v-slot:[channelname]>
        <b>{{ channelname }}</b>
      </template>
    </navigation>

    <div class="toolbar" align="center">
      <!-- <div class="routeback"> -->
        <!-- 返回键 -->
        <el-button class="square-button" @click="routeBack">
          <img src="../assets/icons/back.svg" alt="返回">
        </el-button>
      <!-- </div> -->
      <div class="favorite"  align="center">
        <!-- 收藏键 -->
        <el-button class="square-button" @click="toggleFavorite">
          <i :class="isFavorite ? 'el-icon-star-on' : 'el-icon-star-off'"></i>
        </el-button>
      </div>
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
    <div class="comment-function-area">
      <div class="title">
        <span>
          <el-button type="text" @click="toggleComments">
            {{ commentsVisible ? '收起评论<<<' : '展开评论>>>' }} </el-button>
        </span>
      </div>
      <hr>
      <div v-if="commentsVisible">
        <div v-for="comment in comments" :key="comment.id" class="comment">
          <div class="comment-name">{{ comment.author }}:</div>
          <div class="comment-content">{{ comment.content }}</div>
          <div class="comment-time">发表于 {{ comment.time }}</div>
          <div class="comment-actions">
            <el-button type="text" icon="el-icon-thumb" @click="likes(comment)">
              {{ comment.likes }} 赞
            </el-button>
          </div>

          <div class="comment-reply-field">
            <div v-if="comment.showReplyInput" class="reply-input">
              <el-input v-model="comment.replyContent" placeholder="回复评论" ref="replyInput" @blur="hideReplyInput(comment)"
                v-if="comment.showReplyInput"></el-input>
              <el-button type="primary" @click="reply(comment)">回复</el-button>
            </div>
            <div v-else class="reply-button">
              <el-button type="text" @click="showReplyInput(comment)">
                回复
              </el-button>
            </div>
          </div>

          <div v-for="reply in comment.replies" :key="reply.id" class="comment-reply">
            <div class="comment-reply-name">{{ reply.author }}:</div>
            <div class="comment-reply-content">{{ reply.content }}</div>
            <div class="comment-time">发表于 {{ reply.time }}</div>
            <div class="comment-actions">
              <el-button type="text" icon="el-icon-thumb" @click="likes(reply)">
                {{ reply.likes }} 赞
              </el-button>
            </div>
          </div>
        </div>

        <div class="comment-reply-field">
          <div class="comment-input">
            <el-input v-model="commentContent" placeholder="发表评论" ref="commentInput" @blur="hideCommentInput"></el-input>
          </div>
          <div class="comment-button">
            <el-button type="info" @click="addComment" plain>评论</el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import navigation from "../components/layout/nav.vue"
import { getNews } from '@/api/news';
import backtotop from '../components/layout/backtotop.vue'
import axios from "axios";
import { parentComment, childComment, getCommentList, like ,saveFavoriteStatus, loadFavoriteStatus} from "../api/comments"
import { mapState } from 'vuex';
import { Message } from "element-ui";
export default {
  components: {
    navigation,
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
        { name: '教育', url: '/education' }
      ],
      channel: '',
      pathFromUrl: '',
      news: {
        title: '',
        content: '',
        src: '',
        time: '',
        pic: ''
      },
      comments: [
        {
          id: "1",
          author: "Thomas",
          content: "沙发！",
          likes: 5,
          time: "2023-07-03 16:24:54",
          replies: [
            {
              id: "11",
              author: "John",
              content: "棒",
              likes: 2,
              time: "2023-07-04 16:34:54",
            },
            {
              id: "12",
              author: "用户123",
              content: "厉害",
              likes: 1,
              time: "2023-07-05 23:24:54",
            }
          ]
        },
        {
          id: "2",
          author: "Lisa",
          content: "新闻不错！",
          likes: 3,
          replies: [],
          time: "2023-07-05 22:42:55",
        },
      ],
      isFavorite: false,
      toolbarTop: 0,
      commentContent: "",
      replyContent: "",
      commentsVisible: true
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
    ...mapState('user', ['currentUser']),
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.pathFromUrl = from.fullPath;
      vm.channel = to.params.channel;
    });
  },
  methods: {
    routeBack() {
      // 返回上一页
      this.$router.push({ path: this.pathFromUrl });
      window.close();
    },

    getComments() {
      axios.get("/comments").then((res) => {
        if (Array.isArray(res.data)) {
          this.comments = res.data.map((comment) => ({
            ...comment,
            showReplyInput: false,
            replyContent: "",
          }));
        } else {
          // 处理错误情况，比如服务器返回的数据不是数组
          console.error("Invalid response data:", res.data);
        }
      }).catch((error) => {
        // 处理请求错误
        console.error("Failed to get comments:", error);
      });
    },

    async addComment() {
      if (!this.currentUser.token) {
        Message.warning("请登录后再评论！");
        return;
      }
      if (this.commentContent.trim() === "") {
        Message.warning("请输入评论内容！");
        return;
      }
      const comment = {
        title: this.$route.params.title,
        author: this.currentUser.username,
        content: this.commentContent,
      };
      console.log(comment);
      const res = await parentComment(comment, this.currentUser.token);
      console.log(res);
      const newComment = {
        ...res.data,
        showReplyInput: false,
        replyContent: "",
      };
      this.comments = [
        ...this.comments,
        newComment,
      ];
      this.commentContent = "";
    },

    async reply(comment) {
      if (!this.currentUser) {
        Message.warning("请登录后再评论！");
        return;
      }
      if (this.replyContent.trim() === "") {
        Message.warning("请输入评论内容！");
        return;
      }
      const reply = {
        parentId: comment.id,
        title: this.$route.params.title,
        author: this.currentUser.username,
        content: this.commentContent,
      };
      const res = await childComment(reply, this.currentUser.token);
      const newReply = {
        ...res.data,
      };
      this.comments = [
        ...this.comments.replies,
        newReply,
      ];
      comment.replyContent = "";
      comment.showReplyInput = false;
    },

    async likes(comment) {
      if (this.currentUser.username == "") {
        Message.warning("请登录后再点赞！");
        return;
      }
    },

    async loadFavoriteStatusFromDatabase() {
      try {
        const result = await loadFavoriteStatus(this.news.title);
        this.isFavorite = result.isFavorite;
      } catch (error) {
        console.error('Failed to load favorite status:', error);
      }
    },

    async saveFavoriteStatusToDatabase() {
      try {
        await saveFavoriteStatus(this.news.title, this.isFavorite);
      } catch (error) {
        console.error('Failed to save favorite status:', error);
      }
    },

    async likes(comment) {
      if (this.currentUser.username == "") {
        Message.warning("请登录后再点赞！");
        return;
      }
      const likes = {
        title: this.$route.params.title,
        author: comment.author,
        username: this.currentUser.username,
        content: comment.content,
        commentId: comment.id
      };
      console.log(likes);
      const res = await like(likes, this.currentUser.token)
      console.log(res.msg);
      this.getComments();
    },

    showReplyInput(comment) {
      comment.showReplyInput = true;
      this.$nextTick(() => {
        this.$refs.replyInput.$el.focus();
      });
    },

    hideReplyInput(comment) {
      comment.showReplyInput = false;
    },

    hideCommentInput() {
      this.$refs.commentInput.$el.blur();
    },

    toggleComments() {
      this.commentsVisible = !this.commentsVisible;
    },

    toggleFavorite() {
      // 切换收藏状态
      this.isFavorite = !this.isFavorite;
      this.saveFavoriteStatusToDatabase();
    },
  },
  mounted() {
    this.handleScroll(); // 初始化位置
    this.loadFavoriteStatusFromDatabase();
  },
  created() {
    this.getComments();
    this.$store.dispatch('user/loadCurrentUser');
  }
};
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

.comment-function-area {
  border: 3px solid black;
  background-color: rgb(243, 250, 251);
  border-radius: 3%;
  margin: auto;
  width: 60%;
  padding: 20px;
  font-family: 'sihan-regular', sans-serif;
  margin-bottom: 50px;
}

.comment {
  width: 95%;
  border: 1px solid #cfbfbf;
  border-radius: 20px;
  padding: 15px;
  margin: 10px 0;
  text-align: left;
}

.comment-name {
  font-weight: bold;
  color: #f91c1c;
}

.comment-content {
  margin-top: 10px;
  margin-bottom: 10px;
  color: #333;
}

.comment-reply {
  margin-left: 20px;
  border-left: 1px solid #ccc;
  padding-left: 10px;
}

.comment-reply-field {
  width: 49%;
  margin-left: 20px;
  margin-right: 20px;
}

.comment-reply-name {
  font-weight: bold;
  color: #f91c1c;
}

.comment-reply-content {
  margin-top: 10px;
  margin-bottom: 10px;
  color: #333;
}

.comment-actions {
  margin-top: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.comment-actions .el-button {
  padding-left: 0;
  padding-right: 0;
  color: #666;
}

.comment-actions .el-icon-thumb {
  margin-right: 5px;
  color: #999;
}

.comment-actions .el-icon-message {
  margin-right: 5px;
  color: #999;
}

.reply-input {
  margin-bottom: 10px;
}

.comment-input {
  margin-bottom: 10px;
}

.comment-reply-field {
  display: flex;
  align-items: center;
}

.comment-input {
  flex: 1;
  margin-right: 10px;
}

.comment-button {
  margin-top: -10px;
}

.comment-time {
  color: #666;
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 12px;
}

.toolbar {
  display: block;
  justify-content: center;
  position: fixed;
  top: 50%;
  left: 3%;
  transform: translate(-50%, -50%);
  transition: top 0.3s;
}

.square-button {
  width: 40px;
  height: 40px;
  border-radius: 10%;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center; 
  background-color: #fff;
}

.square-button img {
  height: 25px;
  width: 25px;
}

.favorite {
  margin-top: 2px;
}

.favorite i {
  font-size: 24px;
  color: #f0ad4e;
}
</style>
