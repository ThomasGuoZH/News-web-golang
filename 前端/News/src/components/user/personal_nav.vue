<template>
  <div>
    <el-row class="tac">
      <el-col>
        <div class="t_title">
          <p>个人中心</p>
        </div>
        <el-menu :default-active="$route.path"  class="el-menu" @open="handleOpen" @close="handleClose" :router="true">
          <el-menu-item index="/personal_center/myInfomation">
            <i class="el-icon-user"></i>
            <span slot="title" class="item">我的信息</span>
          </el-menu-item>
          <el-menu-item index="/personal_center/changepassword">
            <i class="el-icon-menu"></i>
            <span slot="title" class="item">更改密码</span>
          </el-menu-item>
            <el-menu-item index="/personal_center/MyMessage/mycomment">
              <i class="el-icon-chat-dot-round"></i>
              <span slot="title" class="item">我的评论</span>
            </el-menu-item>
            <el-menu-item index="/personal_center/MyMessage/myreply">
              <i class="el-icon-chat-dot-square"></i>
              <span slot="title" class="item">回复我的</span>
            </el-menu-item>
            <el-menu-item index="/personal_center/MyMessage/likes">
              <i class="el-icon-coin"></i>
              <span slot="title" class="item">收到的赞</span>
            </el-menu-item>
          <el-menu-item index="/personal_center/MyLikes">
            <i class="el-icon-star-off"></i>
            <span slot="title" class="item">我的收藏</span>
          </el-menu-item>
        </el-menu>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);

    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);

    },
    beforeRouteEnter(to, from, next) {
      next(vm => {
        vm.pathFromUrl = from.fullPath;
        vm.channel = to.params.channel;
      })
    },
  },
  data() {
    return {

    };
  }

}
</script>

<style scoped>

.t_title {
  height: 60px;
  text-align: center;
}

.t_title p {
  padding: 20px 0;
  font-size: 18px;
}
.item{
 font-size: 16px;
}
</style>