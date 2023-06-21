<template>
  <div>
    <!-- <div class="routeback">
      <el-button class='back' @click="routeBack">
        <img src="../../assets/icons/back.svg" alt="返回">
      </el-button>
    </div> -->
    <el-row class="tac">
      <el-col>
        <div class="t_title">
          <p>个人中心</p>
        </div>
        <el-menu default-active="$route.path" class="el-menu" @open="handleOpen" @close="handleClose" :router="true">
          <el-menu-item index="/personal_center/myMessage">
            <i class="el-icon-user"></i>
            <span slot="title">我的信息</span>
          </el-menu-item>
          <el-menu-item index="/personal_center/changepassword">
            <i class="el-icon-menu"></i>
            <span slot="title">更改密码</span>
          </el-menu-item>
          <el-menu-item index="3">
            <i class="el-icon-document"></i>
            <span slot="title">导航三</span>
          </el-menu-item>
          <el-menu-item index="4">
            <i class="el-icon-setting"></i>
            <span slot="title">导航四</span>
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
  }

}
</script>

<style>
.routeback {
  margin-top: 30px;
  margin-left: 15px;
}

.back {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.back img {
  height: 25px;
  width: 25px;
}

.t_title {
  height: 60px;
  text-align: center;
}

.t_title p {
  padding: 20px 0;
}
</style>