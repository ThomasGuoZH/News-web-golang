<template>
  <div class="personal-center">
    <el-container>
      <el-header>基本信息</el-header>
      <el-main>
        <el-form :model="form" :rules="rules" ref="form" label-width="100px">
          <el-form-item label="账户ID">
            <el-textarea v-model="form.id">{{ form.id }}</el-textarea>
          </el-form-item>
          <el-form-item label="姓名" prop="name">
            <el-input v-model="form.username"></el-input>
          </el-form-item>
          <el-form-item label="性别" prop="sex">
            <el-radio-group v-model="form.sex">
              <el-radio label="man">男</el-radio>
              <el-radio label="woman">女</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email"></el-input>
          </el-form-item>
          <el-form-item label="电话" prop="phone">
            <el-input v-model="form.phone"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">保存</el-button>
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>
  
<script>
import { mapState } from 'vuex'
export default {
  data() {
    return {
      form: {
        id: '',
        username: '',
        sex: '',
        email: '',
        phone: '',
      },
    }
  },
  computed: {
    ...mapState('user', ['currentUser']),
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
        confirmPassword: [
          { required: true, message: '请确认密码', trigger: 'blur' },
          { validator: this.validateConfirmPassword, trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
        ],
        sex: [
          { required: true, message: '请选择性别', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '请输入电话号码', trigger: 'blur' },
          { min: 11, max: 11, message: '电话号码必须为11位', trigger: 'blur' }
        ]
      }
      return rules;
    },
  },
  methods: {
    onSubmit() {
      this.$refs.form.validate(valid => {
        if (valid) {
          // 提交表单数据
          this.$store.dispatch('user/changeInfo', {
            form: this.form,
            token: this.currentUser.token,
          })
        } else {
          return false
        }
      })
    }
  },
  created() {
    this.$store.dispatch('user/loadCurrentUser');
    this.form = {
      id: this.currentUser.id,
      username: this.currentUser.username,
      sex: this.currentUser.sex,
      email: this.currentUser.email,
      phone: this.currentUser.phone,
    };
  },
}
</script>
<style>
.personal-center {
  margin: auto;
  width: 900px;
}

.el-header {
  background-color: #B3C0D1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 160px;
}

body>.el-container {
  margin-bottom: 40px;
}
</style>
  