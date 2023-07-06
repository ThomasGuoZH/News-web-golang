<template>
  <div class="MyLikes">
    <el-container>
      <el-header style="font-size: 18px;">我的收藏</el-header>
      <el-main>
        <div class="main_stage">
          <div v-for="(fave, index) in faves" :key="fave.id"
            :class="{ 'likes-Item': true, 'last-Item': index === faves.length - 1 }">
            <el-row>
              <el-col :span="16">
                <router-link :to="'/' + fave.channel + '/newspage/' + fave.title" class="link">
                  <div class="likes-title">{{ fave.title }}</div>
                </router-link></el-col>
              <el-col :span="4" :offset="3">
                <div class="likes-time">{{ fave.time }}</div>
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
import { getFavsList } from '@/api/favourite';
export default {
  data() {
    return {
      faves: []
    };
  },
  computed: {
    ...mapState('user', ['currentUser']),
  },
  methods: {
    async getFavsList() {
      const res = await getFavsList(this.currentUser.id);
      console.log(res);
      if (res.code === 200 && Array.isArray(res.data.favourites)) {
        this.faves = res.data.favourites.map((fave) => ({
          ...fave,
        }));
      }
      console.log(this.faves);
    }
  },
  created() {
    this.getFavsList();
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
  white-space: nowrap;
  height: 50px;
  text-align: left;
}
</style>
