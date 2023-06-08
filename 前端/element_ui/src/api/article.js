import { articleAPI } from './request.js'


const getArticleList = articleAPI('/articleapi')

export {
    getArticleList,
}