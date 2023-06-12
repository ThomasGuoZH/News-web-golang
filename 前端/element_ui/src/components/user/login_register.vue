<template>
  <div class="login_register">
    <el-button @click="dialogFormVisible = true" class="button_style" type="primary" round>登录</el-button>
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
            <el-form-item v-show="!isLogin" label="性别" prop="sex">
              <el-radio-group v-model="form.sex">
                <el-radio label="man">男</el-radio>
                <el-radio label="woman">女</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item v-show="!isLogin" label="电话号码" prop="phone">
              <el-input v-model="form.phone"></el-input>
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
import { userRegister, userLogin } from '@/api/user.js'
import { Message } from 'element-ui'
export default {
  data() {
    return {
      registerResponse: '',
      loginResponse: '',
      dialogFormVisible: false,
      isLogin: true,
      dTitle: '登录',
      form: {
        username: '',
        password: '',
        confirmPassword: '',
        email: '',
        sex: '',
        phone: '',
      },
      loginForm: {
        username: '',
        password: '',
      }
    }
  },
  computed: {
    rules() {
      const rules = {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 2, max: 10, message: '用户名在2到10个字符之间', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 16, message: '密码在6到16个字符之间', trigger: 'blur' }
        ],
        confirmPassword: this.isLogin ? [] : [
          { required: true, message: '请确认密码', trigger: 'blur' },
          { validator: this.validateConfirmPassword, trigger: 'blur' }
        ],
        email: this.isLogin ? [] : [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
        ],
        sex: this.isLogin ? [] : [
          { required: true, message: '请选择性别', trigger: 'blur' }
        ],
        phone: this.isLogin ? [] : [
          { required: true, message: '请输入电话号码', trigger: 'blur' },
          { min: 11, max: 11, message: '电话号码必须为11位', trigger: 'blur' }
        ]
      };
      return rules;
    }
  },
  methods: {
    async submitForm() {
      const valid = await this.$refs.form.validate();
      if (valid) {
        if (this.isLogin) {
          // 发送登录请求
          console.log(this.isLogin)
          this.loginForm.username = this.form.username;
          this.loginForm.password = this.form.password;
          console.log(this.loginForm);
          this.loginResponse = await userLogin(this.loginForm);
          if (this.loginResponse.msg === '登录成功!') {
            Message.success(this.loginResponse.msg);
          } else if (this.loginResponse.msg === '用户不存在!') {
            Message.warning(this.loginResponse.msg)
          } else if (this.loginResponse.msg === '密码错误!') {
            Message.warning(this.loginResponse.msg)
          }
        } else {
          // 发送注册请求
          console.log(this.form.username);
          console.log(this.form.password);
          this.registerResponse = await userRegister(this.form);
          if (this.registerResponse === '注册成功!') {
            Message.success(this.registerResponse);
          } else if (this.registerResponse === '用户名已被使用!') {
            Message.warning(this.registerResponse)
          }
        }
      }
      else {
        return false
      }
    },
    validateConfirmPassword(rule, value, callback) {
      if (!this.isLogin) {
        if (value !== this.form.password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
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