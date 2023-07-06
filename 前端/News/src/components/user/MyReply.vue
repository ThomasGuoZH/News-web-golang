<template>
  <div class="MyReply">
    <el-container>
      <el-header style="font-size: 18px;">回复我的</el-header>
      <el-main>
        <div class="main_stage">
          <div v-for="(item, index) in replies" :key="item.id"
            :class="{ 'reply-Item': true, 'last-Item': index === replies.length - 1 }">
            <el-row>
              <el-col :span="16">
                <div class="reply-name"><strong>{{ item.author }}</strong>回复了我的评论</div>
                <router-link :to="'/' + item.channel + '/newspage/' + item.title" class="link">
                  <div class="relpy-othercomment">{{ item.content }}</div>
                </router-link>
                <div class="reply-time">{{ item.time }}</div>
              </el-col>
              <el-col :span="6" :offset="1">
                <router-link :to="'/' + item.channel + '/newspage/' + item.title" class="link">
                  <div class="reply-comment">{{ item.parentContent }}</div>
                </router-link>
              </el-col>
            </el-row>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import { getRepliesList } from '@/api/personal';
export default {
  data() {
    return {
      replies: []
    };
  },
  computed: {
    ...mapState('user', ['currentUser']),
  },
  methods: {
    async getRepliesList() {
      const res = await getRepliesList(this.currentUser.id);
      console.log(res);
      if (res.code === 200 && Array.isArray(res.data.replies)) {
        this.replies = res.data.likes.map((reply) => ({
          ...reply,
        }));
      }
      console.log(this.replies);
    }
  },
  created() {
    this.getRepliesList();
    this.$store.dispatch('user/loadCurrentUser');
  }
}
</script>

<style>
.main_stage {
  background-color: rgb(255, 255, 255);
  border-radius: 10px;
}

.link {
  text-decoration: none;
  color: black
}

.reply-Item {
  display: block;
  line-height: 30px;
  border-bottom: 1px solid rgba(77, 75, 75, 0.5);
  height: auto;
}

.last-Item {
  border-bottom: none;
}

.reply-name {
  height: 40px;
  text-align: left;
  margin-left: 10px;
  font-size: 15px;
  margin-top: 5px;
}

.relpy-othercomment {
  height: auto;
  text-align: left;
  margin-left: 10px;
}

.reply-comment {
  line-height: 105px;
  height: 105px;
  text-align: left;
  margin-left: 80px;
  color: rgba(152, 171, 171, 0.7);
}

.reply-time {
  font-size: small;
  height: auto;
  text-align: left;
  margin-left: 10px;
}
</style>
