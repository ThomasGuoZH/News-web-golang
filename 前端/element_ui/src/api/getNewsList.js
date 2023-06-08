import { newsAPI } from './request.js'


const getNewsList = newsAPI('/newsapi')

export {
    getNewsList,
}