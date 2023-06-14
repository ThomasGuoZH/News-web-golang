<template>
  <div class="personal-center">
    <el-container>
      <el-header>基本信息</el-header>
      <el-main>

        <el-form :model="form" :rules="rules" ref="form" label-width="100px">
          <el-form-item label="账户ID" prop="id">
            <el-input v-model="form.name" :value="name"></el-input>
          </el-form-item>
          <el-form-item label="姓名" prop="name">
            <el-input v-model="form.name" :value="name"></el-input>
          </el-form-item>
          <el-form-item label="性别" prop="sex">
            <el-radio-group v-model="form.sex" :value="sex">
              <el-radio label="man">男</el-radio>
              <el-radio label="woman">女</el-radio>
              <el-radio label="secret">保密</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email" :value="email"></el-input>
          </el-form-item>
          <el-form-item label="电话" prop="phone">
            <el-input v-model="form.phone" :value="phone"></el-input>
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
export default {
  data() {
    return {
      form: {
        id: '',
        name: '',
        sex: '',
        email: '',
        phone: '',
      }
    }
  },
  methods: {
    onSubmit() {
      this.$refs.form.validate(valid => {
        if (valid) {
          // 提交表单数据
          this.$emit('submit', this.form)
        } else {
          return false
        }
      })
    }
  }
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
  