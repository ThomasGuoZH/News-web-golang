import { getPersonalListAPI } from "./request";

//获取个人主页我的评论列表
const getPersonalCommentList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/comments')
//获取个人主页收到的赞列表
const getLikesList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/likes')
//获取个人主页回复我的列表
const getRepliesList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/replies')

export {
    getPersonalCommentList,
    getRepliesList,
    getLikesList
}