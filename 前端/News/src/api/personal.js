import { getPersonalListAPI } from "./request";

const getPersonalCommentList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/comments')
const getLikesList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/likes')
const getRepliesList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/replies')

export {
    getPersonalCommentList,
    getRepliesList,
    getLikesList
}