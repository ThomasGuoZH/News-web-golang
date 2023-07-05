import { getPersonalCommentsListAPI, getLikesListAPI, getRepliesListAPI } from "./request";

const getPersonalCommentList = getPersonalCommentsListAPI('http://127.0.0.1:8080/api/personal/comments')
const getLikesList = getLikesListAPI('http://127.0.0.1:8080/api/personal/likes')
const getRepliesList = getRepliesListAPI('http://127.0.0.1:8080/api/personal/replies')

export {
    getPersonalCommentList,
    getRepliesList,
    getLikesList
}