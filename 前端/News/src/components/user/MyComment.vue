<template>
  <div class="MyComment">
    <el-container>
      <el-header>我的评论</el-header>
      <el-main>
        <div class="main_stage">
          <div v-for="(comment, index) in comments" :key="comment.id"
            :class="{ 'comment-Item': true, 'last-Item': index === comments.length - 1 }">
            <router-link :to="'/' + comment.channel + '/newspage/' + comment.title" class="link">
              <div class="comment-title">{{ comment.title }}</div>
              <div class="comment-context">{{ comment.content }}</div>
            </router-link>
            <div class="comment-time">{{ comment.time }}</div>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import { getPersonalCommentList } from '@/api/personal';
export default {
  data() {
    return {
      comments: []
    };
  },
  computed: {
    ...mapState('user', ['currentUser']),
  },
  methods: {
    async getComments() {
      const user_id = this.currentUser.id
      const res = await getPersonalCommentList(user_id);
      console.log(res);
      if (res.code === 200 && Array.isArray(res.data.comments)) {
        this.comments = res.data.comments.map((comment) => ({
          ...comment,
        }));
      }
    },
  },
  created() {
    this.$store.dispatch('user/loadCurrentUser');
    this.getComments();
  }
}
</script>

<style>
.main_stage {
  background-color: rgb(255, 255, 255);
  border-radius: 10px;
}

.comment-Item {
  display: block;
  height: 120px;
  line-height: 50px;
  border-bottom: 1px solid rgba(77, 75, 75, 0.5);
}

.last-Item {
  border-bottom: none;
}

.comment-title {
  font-size: medium;
  height: 40px;
  text-align: left;
  margin-left: 5px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.comment-context {
  font-size: large;
  height: 40px;
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-left: 10px;
}

.comment-time {
  font-size: small;
  height: auto;
  text-align: left;
  margin-left: 30px;
}
</style>