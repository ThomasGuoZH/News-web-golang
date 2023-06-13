<template>
    <div class="comment-function-area">
        <div class="title">          
            <span>
                <p>评论</p>
                <el-button type="text" @click="toggleComments">
                {{ commentsVisible ? '收起评论<<<' : '展开评论>>>' }}
                </el-button>
            </span>
        </div>
        <hr>
        <div v-if="commentsVisible">
        <div v-for="comment in comments" :key="comment.id" class="comment">
            <div class="comment-name">{{ comment.name }}:</div>
            <div class="comment-content">{{ comment.content }}</div>
            <div class="comment-actions">
            <el-button type="text" icon="el-icon-thumb" @click="like(comment)">
                {{ comment.likes }} 赞
            </el-button>
            </div>
            
            <div class="comment-reply-field">
                <div v-if="comment.showReplyInput" class="reply-input">
                    <el-input
                    v-model="comment.replyContent"
                    placeholder="回复评论"
                    ref="replyInput"
                    @blur="hideReplyInput(comment)"
                    ></el-input>
                </div>
                <div v-else class="reply-button">
                    <el-button type="primary" @click="showReplyInput(comment)">
                        回复
                    </el-button>
                </div>
            </div>

            <div v-for="reply in comment.replies" :key="reply.id" class="comment-reply">
            <div class="comment-reply-name">{{ reply.name }}:</div>
            <div class="comment-reply-content">{{ reply.content }}</div>
            <div class="comment-actions">
                <el-button type="text" icon="el-icon-thumb" @click="like(reply)">
                {{ reply.likes }} 赞
                </el-button>
            </div>
            </div>
        </div>
    
        <div class="comment-reply-field">
            <div class="comment-input">
            <el-input
                v-model="commentContent"
                placeholder="发表评论"
                ref="commentInput"
                @blur="hideCommentInput"
            ></el-input>
            </div>
            <div class="comment-button">
            <el-button type="primary" @click="addComment">评论</el-button>
            </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  import store from "../../store/index.js";
  
  export default {
    data() {
    return {
        comments: [
        {
            id: 1,
            name: "用户1",
            content: "这是第一条评论",
            likes: 5,
            replies: [
            {
                id: 11,
                name: "用户2",
                content: "回复1",
                likes: 2
            },
            {
                id: 12,
                name: "用户3",
                content: "回复2",
                likes: 1
            }
            ]
        },
        {
            id: 2,
            name: "用户4",
            content: "这是第二条评论",
            likes: 3,
            replies: []
        }
        ],
        commentContent: "",
        replyContent: "",
        commentsVisible: true
    };
    },

    created() {
      this.getComments();
    },
    methods: {
        getComments() {
      axios.get("/comments").then((res) => {
        if (Array.isArray(res.data)) {
          this.comments = res.data.map((comment) => ({
            ...comment,
            showReplyInput: false,
            replyContent: "",
          }));
        } else {
          // 处理错误情况，比如服务器返回的数据不是数组
          console.error("Invalid response data:", res.data);
        }
      }).catch((error) => {
        // 处理请求错误
        console.error("Failed to get comments:", error);
      });
    },
      addComment() {
        const comment = {
          name: "匿名用户",
          content: this.commentContent,
        };
        axios.post("/comments", comment).then((res) => {
          this.comments = res.data.map((comment) => ({
            ...comment,
            showReplyInput: false,
            replyContent: "",
          }));
          this.commentContent = "";
        });
      },
      reply(comment) {
        const reply = {
          name: "匿名用户",
          content: comment.replyContent,
        };
        axios
          .post(`/comments/${comment.id}/replies`, reply)
          .then((res) => {
            comment.replies = res.data;
            comment.replyContent = "";
            comment.showReplyInput = false;
          });
      },
      like(item) {
        store.commit("addLike", item.id);
        axios.put(`/comments/${item.id}`, { likes: store.state.likes[item.id] }).then((res) => {
          store.commit("setLikes", { id: item.id, likes: res.data.likes });
        });
      },
      showReplyInput(comment) {
        comment.showReplyInput = true;
        this.$nextTick(() => {
          this.$refs.replyInput.$el.focus();
        });
      },
      hideReplyInput(comment) {
        comment.showReplyInput = false;
      },
      hideCommentInput() {
        this.$refs.commentInput.$el.blur();
      },
      toggleComments() {
      this.commentsVisible = !this.commentsVisible;
    },
    },
  };
  </script>
  
  <style>
  .comment-function-area {
     border: 1px solid #333;
     margin-left: 20px;
     margin-right: 20px;
     margin-bottom: 20px;
     width: 80%;
     padding: 20px;
     background-color: rgb(252, 253, 254);
  }

  .comment {
    width: 90%;
    border: 1px solid #cfbfbf;
    border-radius: 20px;
    padding: 15px;
    margin-top: 10px;
    margin-bottom: 10px;
   text-align: left;
  }
  
  .comment-name {
    font-weight: bold;
    color: #333;
  }
  
  .comment-content {
    margin-top: 10px;
    margin-bottom: 10px;
    color: #666;
  }
  
  .comment-reply {
    margin-left: 20px;
    border-left: 1px solid #ccc;
    padding-left: 10px;
  }
  
  .comment-reply-field {
    width: 49%;
    margin-left: 20px;
    margin-right: 20px;
  }
  
  .comment-reply-name {
    font-weight: bold;
    color: #333;
  }
  
  .comment-reply-content {
    margin-top: 10px;
    margin-bottom: 10px;
    color: #666;
  }
  
  .comment-actions {
    margin-top: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .comment-actions .el-button {
    padding-left: 0;
    padding-right: 0;
    color: #666;
  }
  
  .comment-actions .el-icon-thumb {
    margin-right: 5px;
    color: #999;
  }
  
  .comment-actions .el-icon-message {
    margin-right: 5px;
    color: #999;
  }
  
  .reply-input {
    margin-bottom: 10px;
  }
  
  .comment-input {
    margin-bottom: 10px;
  }
  </style>
  