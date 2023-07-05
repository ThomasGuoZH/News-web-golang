<template>
  <div class="MyComment">
    <el-container>
      <el-header>我的评论</el-header>
      <el-main>
        <div class="main_stage">
        <div v-for="(item,index) in comment" :key="item.id" :class="{'comment-Item': true, 'last-Item': index === comment.length - 1}">
          <router-link :to="'/politics/newspage/' + item.title" class="link">
            <div class="comment-title">{{ item.title }}</div>
            <div class="comment-context">{{ item.context}}</div>
            </router-link>
            <div class="comment-time">{{ item.time }}</div>
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
            comment:[
              {
                id: 1,
                title:'这是新闻标题',
                context:'这里是所发评论',
                time:'2023/07/04',
            },
            {
              id:2,
              title:"02",
              context:"0.0",
              time:"2023-07",
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
}
.comment-Item{
  display:block;
  height:120px;
  line-height: 50px;
  border-bottom:1px solid rgba(77, 75, 75, 0.5);
}
.last-Item{
  border-bottom:none;
}
.comment-title{
  font-size: medium;
  height:40px;
  text-align: left;
  margin-left: 5px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.comment-context{
  font-size: large;
  height:40px;
  text-align: left;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-left:10px;
}
.comment-time{
  font-size: small;
  height:auto;
  text-align: left;
  margin-left: 30px;
}
</style>