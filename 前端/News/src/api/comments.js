import { parentCommentAPI, childCommentAPI, getCommentsListAPI, LikesAPI } from "./request";

const parentComment = parentCommentAPI('http://127.0.0.1:8080/api/comment/parent_comment')
const childComment = childCommentAPI('http://127.0.0.1:8080/api/comment/child_comment')
const getCommentList = getCommentsListAPI('http://127.0.0.1:8080/api/comment/comment_list')
const like = LikesAPI('http://127.0.0.1:8080/api/comment/likes')

export {
    parentComment,
    childComment,
    getCommentList,
    like
}