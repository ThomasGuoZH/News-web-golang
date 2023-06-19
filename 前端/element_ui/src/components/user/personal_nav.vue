<template>
  <div>
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
    }
  }

}
</script>

<style>
.t_title {
  height: 60px;
  text-align: center;
}

.t_title p {
  padding: 20px 0;
}
</style>