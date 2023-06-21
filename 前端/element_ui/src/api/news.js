import { storeNewsAPI, getNewsAPI } from "./request";

const storeNews = storeNewsAPI('http://127.0.0.1:8080/api/news/store_news')
const getNews = getNewsAPI('http://127.0.0.1:8080/api/news/get_news')

export {
    storeNews,
    getNews
}