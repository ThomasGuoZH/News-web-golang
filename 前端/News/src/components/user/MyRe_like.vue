<template>
  <div class="MyRelike">
    <el-container>
      <el-header>收到的赞</el-header>
      <el-main>
        <div class="main_stage">
        <div v-for="(item,index) in rlikes" :key="item.id" :class="{'rlikes-Item': true, 'last-Item': index === rlikes.length - 1}">
          <el-row>
            <el-col :span="16">
              <div class="rlikes-name"><strong>{{ item.name }}</strong>赞了我的评论</div> 
              <div class="rlikes-time">{{ item.time }}</div>
            </el-col>
            <el-col :span="6" :offset="1">
              <router-link :to="'/politics/newspage/' + item.title" class="link">
              <div class="rlikes-comment">{{ item.mycomment }}</div>
              </router-link>
            </el-col>
          </el-row>
        </div>
      </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
    data(){
        return{
            rlikes:[
              {
                id: 1,
                name:'点赞用户',
                mycomment:"这是被点赞的评论",
                time:'2023/07/04',
                tiele:''
            },
            {
              id:2,
              name:"02",
              mycomment:"doge",
              time:"2023-07",
              title:''
            }
          ]
        };
    }
}
</script>

<style>
.main_stage{
  background-color: pink;
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
}
.rlikes-Item{
  display:block;
  line-height: 50px;
  border-bottom:1px solid rgba(77, 75, 75, 0.5);
  height:auto;
}
.last-Item{
  border-bottom:none;
}
.rlikes-name{
  height:50px;
  text-align: left;
  margin-left: 10px;
}
.rlikes-comment{
  height:auto;
  text-align: left;
  margin-left: 10px;
  color:rgba(152, 171, 171, 0.7);
}
.rlikes-time{
  font-size: small;
  height:auto;
  text-align: left;
  margin-left: 30px;
}
</style>
