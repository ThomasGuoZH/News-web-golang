<template>
    <div>
        
        <el-button type="primary" @click="dialogFormVisible = true">登录</el-button>
        
        <el-dialog title="登录" :visible.sync="dialogFormVisible" center>   
        <el-form :model="form">
          <el-form-item label="用户名：" :label-width="formLabelWidth" >
          <el-input v-model="form.name" auto-complete="off" placeholder="请输入用户名" clearable></el-input>
          </el-form-item>
          <el-form-item label="密码：" :label-width="formLabelWidth" >
          <el-input v-model="form.password" auto-complete="off" placeholder="请输入密码" clearable show-password></el-input>
          </el-form-item>
              <el-link  class="register" type="primary" @click="change_register">没有账号？立即注册</el-link>
         </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="dialogFormVisible = false">{{ isLogin ? '登录' : '注册' }}</el-button>
    </div>
    </el-dialog>
    </div>
</template>
<script>
export default {
  name: 'login',
  data(){
    // tt:{
    return{
        dialogFormVisible: false,
        form: {
          name: '',
          password: ''
        },
        formLabelWidth: '120px',
        isLogin:'false'
    }
  },
  methods:{
      change_register(){

      }
  }
}
</script>
<style scoped>
  .register{
    position: relative;
    left:60px;
  }
</style>