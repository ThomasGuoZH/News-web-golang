<template>
  <div class="change-password">
    <el-container>
      <el-header>修改密码</el-header>
      <el-main>
        <el-form :model="form" :rules="rules" ref="form" label-width="100px" label-position="right">
          <el-col span="12" offset="5">
            <el-form-item label="旧密码" prop="oldPassword">
              <el-input type="password" v-model="form.oldPassword" style="width: 300px"></el-input>
            </el-form-item>

            <el-form-item label="新密码" prop="newPassword">
              <el-input type="password" v-model="form.newPassword" style="width: 300px"></el-input>
            </el-form-item>

            <el-form-item label="确认密码" prop="confirmPassword">
              <el-input type="password" v-model="form.confirmPassword" style="width: 300px"></el-input>
            </el-form-item></el-col>
          <el-col span="1" offset="8">
            <el-form-item>
              <el-button type="primary" @click="onSubmit">保存</el-button>
            </el-form-item></el-col>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>
  
<script>
export default {
  data() {
    return {
      form: {
        oldPassword: '',
        newPassword: '',
        confirmPassword: ''
      },
      rules: {
        oldPassword: [
          { required: true, message: '请输入旧密码', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          { min: 6, message: '密码长度至少为6', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          { validator: this.validateConfirmPassword, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    validateConfirmPassword(rule, value, callback) {
      if (value !== this.form.newPassword) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    },
    onSubmit() {
      this.$refs.form.validate(valid => {
        if (valid) {
          // 提交表单数据
          console.log(this.form)
        } else {
          return false
        }
      })
    }
  }
}
</script>
<style></style>