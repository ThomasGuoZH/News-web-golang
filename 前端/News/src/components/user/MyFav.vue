<template>
  <div class="MyLikes">
    <el-container>
      <el-header>我的收藏</el-header>
      <el-main>
        <div class="main_stage">
          <div v-for="(like, index) in likes" :key="like.id"
            :class="{ 'likes-Item': true, 'last-Item': index === likes.length - 1 }">
            <el-row>
              <el-col :span="16">
                <router-link :to="'/' + like.channel + '/newspage/' + like.title" class="link">
                  <div class="likes-title">{{ like.title }}</div>
                </router-link></el-col>
              <el-col :span="4" :offset="2">
                <div class="likes-time">{{ like.time }}</div>
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
export default {
  data() {
    return {
      likes: []
    };
  },
  computed: {
    ...mapState('user', ['currentUser']),
  },
  methods: {

  },
  created() {
    this.$store.dispatch('user/loadCurrentUser');
  }
}
</script>

<style>
.main_stage {
  background-color: rgb(255, 255, 255);
  border-radius: 10px;
}

.likes-Item {
  display: block;
  height: 50px;
  line-height: 50px;
  border-bottom: 1px solid rgba(77, 75, 75, 0.5);
}

.last-Item {
  border-bottom: none;
}

.likes-title {
  height: 50px;
  text-align: left;
  margin-left: 10px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.likes-time {
  height: 50px;
  text-align: left;
}
</style>
