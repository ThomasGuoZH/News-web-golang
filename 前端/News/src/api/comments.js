import { parentCommentAPI, childCommentAPI, getNewsCommentsListAPI, LikesAPI, getLikesListAPI } from "./request";

const parentComment = parentCommentAPI('http://127.0.0.1:8080/api/comment/parent_comment')
const childComment = childCommentAPI('http://127.0.0.1:8080/api/comment/child_comment')
const getNewsCommentList = getNewsCommentsListAPI('http://127.0.0.1:8080/api/comment/comment_list')
const like = LikesAPI('http://127.0.0.1:8080/api/comment/likes')
const getLikesList = getLikesListAPI('http://127.0.0.1:8080/api/comment/')

export {
    parentComment,
    childComment,
    getNewsCommentList,
    like
}