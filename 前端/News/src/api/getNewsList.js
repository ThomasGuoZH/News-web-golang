import { newsAPI } from './request.js'

//新闻api接口
const getNewsList = newsAPI('/newsapi')

export {
    getNewsList,
}