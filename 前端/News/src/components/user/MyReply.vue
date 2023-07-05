<template>
  <div class="MyReply">
    <el-container>
      <el-header>回复我的</el-header>
      <el-main>
        <div class="main_stage">
          <div v-for="(item, index) in reply" :key="item.id"
            :class="{ 'reply-Item': true, 'last-Item': index === reply.length - 1 }">
            <el-row>
              <el-col :span="16">
                <div class="reply-name"><strong>{{ item.name }}</strong>回复了我的评论</div>
                <router-link :to="'/' + item.channel + '/newspage/' + item.title" class="link">
                  <div class="relpy-othercomment">{{ item.othercommet }}</div>
                </router-link>
                <div class="reply-time">{{ item.time }}</div>
              </el-col>
              <el-col :span="6" :offset="1">
                <router-link :to="'/' + item.channel + '/newspage/' + item.title" class="link">
                  <div class="reply-comment">{{ item.mycomment }}</div>
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
  data() {
    return {
      reply: [
        {
          id: 1,
          name: '点赞用户',
          othercommet: "这是该用户对我评论的评论",
          mycomment: "这是被点赞的评论",
          time: '2023/07/04',
          title: '',
          channel: ''
        },
        {
          id: 2,
          name: "02",
          othercommet: "Oh man",
          mycomment: "doge",
          time: "2023-07",
          title: '',
          channel: ''
        }
      ]
    };
  }
}
</script>

<style>
.main_stage {
  background-color: rgb(255, 255, 255);
  border-radius: 10px;
}

.link {
  text-decoration: none;
  color: black
}

.reply-Item {
  display: block;
  line-height: 50px;
  border-bottom: 1px solid rgba(77, 75, 75, 0.5);
  height: auto;
}

.last-Item {
  border-bottom: none;
}

.reply-name {
  height: 50px;
  text-align: left;
  margin-left: 10px;
}

.relpy-othercomment {
  height: auto;
  text-align: left;
  margin-left: 10px;
}

.reply-comment {
  height: auto;
  text-align: left;
  margin-left: 10px;
  color: rgba(152, 171, 171, 0.7);
}

.reply-time {
  font-size: small;
  height: auto;
  text-align: left;
  margin-left: 30px;
}
</style>
