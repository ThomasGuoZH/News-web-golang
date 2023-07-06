import { getNewsCommentsListAPI, tokenAPI, deleteCommentAPI } from "./request";

//发布一级评论
const parentComment = tokenAPI('http://127.0.0.1:8080/api/comment/parent_comment')
//发布二级评论
const childComment = tokenAPI('http://127.0.0.1:8080/api/comment/child_comment')
//获取新闻业评论列表
const getNewsCommentList = getNewsCommentsListAPI('http://127.0.0.1:8080/api/comment/comment_list')
//评论点赞
const like = tokenAPI('http://127.0.0.1:8080/api/comment/likes')
//判断是否点赞
const isLiked = tokenAPI('http://127.0.0.1:8080/api/comment/is_liked')
//删除一级
const deleteComment = deleteCommentAPI('http://127.0.0.1:8080/api/personal/del_comment')
export {
    parentComment,
    childComment,
    getNewsCommentList,
    like,
    isLiked,
    deleteComment
}