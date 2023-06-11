<template>
    <div class="comment-section">
        <h3>评论区</h3>
        <div class="comment-list">
            <div v-for="comment in comments" :key="comment.id" class="comment">
                <div class="comment-header">
                    <span class="comment-author">{{ comment.author }}</span>
                    <span class="comment-date">{{ comment.date }}</span>
                </div>
                <div class="comment-content">{{ comment.content }}</div>
                <div class="comment-footer">
                    <span class="comment-like" @click="likeComment(comment.id)">
                        <i class="fa fa-thumbs-up"></i>
                        {{ likes[comment.id] || 0 }}
                    </span>
                    <span class="comment-reply" @click="replyComment(comment.id)">
                        <i class="fa fa-reply"></i>
                        回复
                    </span>
                </div>
            </div>
        </div>
        <div class="comment-form">
            <h4>发表评论</h4>
            <form @submit.prevent="submitComment">
                <div class="form-group">
                    <label for="comment-author">昵称</label>
                    <input type="text" id="comment-author" v-model="author" required>
                </div>
                <div class="form-group">
                    <label for="comment-content">评论内容</label>
                    <textarea id="comment-content" v-model="content" required></textarea>
                </div>
                <button type="submit">提交评论</button>
            </form>
        </div>
    </div>
</template>

<!-- <script>
export default {
    data() {
        return {
            author: '',
            content: '',
        };
    },
    computed: {
        comments() {
            return this.$store.state.comments;
        },
        likes() {
            return this.$store.state.likes;
        },
    },
    methods: {
        submitComment() {
            const comment = {
                id: Date.now(),
                author: this.author,
                content: this.content,
                date: new Date().toLocaleString(),
            };
            this.$store.commit('addComment', comment);
            this.author = '';
            this.content = '';
        },
        likeComment(id) {
            this.$store.commit('likeComment', id);
        },
        replyComment(id) {
            // TODO: 实现回复功能
        },
    },
};
</script> -->