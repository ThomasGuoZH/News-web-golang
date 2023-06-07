<template>
  <div class="login_register">
    <el-link @click="dialogFormVisible = true" class="button_style" type="primary">登录</el-link>
    <el-dialog :title="dTitle" :visible.sync="dialogFormVisible" center>
      <el-row class="log">
        <el-col :span="16">
          <el-form :model="form" :rules="rules" ref="form" label-width="80px">

            <el-form-item label="用户名" prop="username">
              <el-input v-model="form.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input type="password" v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item v-show="!isLogin" label="确认密码" prop="confirmPassword">
              <el-input type="password" v-model="form.confirmPassword"></el-input>
            </el-form-item>
            <el-form-item v-show="!isLogin" label="邮箱" prop="email">
              <el-input v-model="form.email"></el-input>
            </el-form-item>
            <el-form-item>
              <el-checkbox v-model="form.remember">记住密码</el-checkbox>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-link type="primary" @click="isLogin = false, dTitle = '注册'" v-show="isLogin">没有账号？立即注册</el-link>
        <el-button type="primary" @click="isLogin = true, dTitle = '登录'" v-show="!isLogin">返回</el-button>
        <el-button class="title" type="primary" @click="submitForm">{{ isLogin ? '登录' : '注册' }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>
  
<script>
export default {
  // data:{
  //   dTitle,
  // },
  data() {
    return {
      dialogFormVisible: false,
      isLogin: true,
      dTitle: '登录',
      form: {
        username: '',
        password: '',
        confirmPassword: '',
        email: '',
        remember: false
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请确认密码', trigger: 'blur' },
          { validator: this.validateConfirmPassword, trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submitForm() {
      this.$refs.form.validate(valid => {
        if (valid) {
          if (this.isLogin) {
            // 发送登录请求
          } else {
            // 发送注册请求
          }
        } else {
          return false
        }
      })
    },
    validateConfirmPassword(rule, value, callback) {
      if (value !== this.form.password) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    },
  }
}
</script>
  
<style >
.log {
  position: relative;
  left: 120px;
}

/* 样式 */
</style>
  