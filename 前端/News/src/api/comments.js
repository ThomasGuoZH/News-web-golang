import { getNewsCommentsListAPI, tokenAPI } from "./request";

const parentComment = tokenAPI('http://127.0.0.1:8080/api/comment/parent_comment')
const childComment = tokenAPI('http://127.0.0.1:8080/api/comment/child_comment')
const getNewsCommentList = getNewsCommentsListAPI('http://127.0.0.1:8080/api/comment/comment_list')
const like = tokenAPI('http://127.0.0.1:8080/api/comment/likes')
const isLiked = tokenAPI('http://127.0.0.1:8080/api/comment/is_liked')

export {
    parentComment,
    childComment,
    getNewsCommentList,
    like,
    isLiked
}