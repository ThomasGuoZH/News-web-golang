import { storeNewsAPI, getNewsAPI } from "./request";

//存储页面中的新闻
const storeNews = storeNewsAPI('http://127.0.0.1:8080/api/news/store_news')
//获取新闻页的新闻
const getNews = getNewsAPI('http://127.0.0.1:8080/api/news/get_news')

export {
    storeNews,
    getNews
}