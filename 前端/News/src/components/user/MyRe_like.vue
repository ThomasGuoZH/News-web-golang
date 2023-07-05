<template>
  <div class="MyRelike">
    <el-container>
      <el-header>收到的赞</el-header>
      <el-main>
        <div class="main_stage">
          <div v-for="(item, index) in rlikes" :key="item.id"
            :class="{ 'rlikes-Item': true, 'last-Item': index === rlikes.length - 1 }">
            <router-link :to="'/' + item.channel + '/newspage/' + item.title" class="link">
              <el-row>
                <el-col :span="16">
                  <div class="rlikes-name"><strong>{{ item.liker }}</strong>赞了我的评论</div>
                  <div class="rlikes-time">{{ item.time }}</div>
                </el-col>
                <el-col :span="6" :offset="1">
                  <div class="rlikes-comment">{{ item.content }}</div>
                </el-col>
              </el-row>
            </router-link>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import { getLikesList } from '@/api/personal';
export default {
  data() {
    return {
      rlikes: []
    };
  },
  computed: {
    ...mapState('user', ['currentUser']),
  },
  methods: {
    async getLikesList() {
      const res = await getLikesList(this.currentUser.id);
      console.log(res);
      if (res.code === 200 && Array.isArray(res.data.likes)) {
        this.rlikes = res.data.likes.map((like) => ({
          ...like,
        }));
      }
      console.log(this.rlikes);
    }
  },
  created() {
    this.getLikesList();
    this.$store.dispatch('user/loadCurrentUser');
  }
}
</script>

<style>
.main_stage {
  background-color: rgb(255, 255, 255);
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
}

.rlikes-Item {
  display: block;
  line-height: 40px;
  border-bottom: 1px solid rgba(77, 75, 75, 0.5);
  height: auto;
}

.last-Item {
  border-bottom: none;
}

.rlikes-name {
  height: 50px;
  text-align: left;
  margin-left: 10px;
}

.rlikes-comment {
  line-height: 90px;
  height: 90px;
  text-align: left;
  margin-left: 10px;
  color: rgba(152, 171, 171, 0.7);
}

.rlikes-time {
  font-size: small;
  height: 20px;
  text-align: left;
  margin-left: 10px;
}
</style>
