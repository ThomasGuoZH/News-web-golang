import { getarticleTitleapi, getarticlePicapi, getarticleContentapi } from './request'

axios.defaults.baseURL = 'https://api.jisuapi.com/news/get?appkey=14b3d45a212c10d0&start=0&num=20';
const getarticleTitle = getarticleTitleapi(baseURL)
const getarticlePic = getarticlePicapi(baseURL)
const getarticleContent = getarticleContentapi(baseURL)

export {
    getarticleTitle,
    getarticlePic,
    getarticleContent
}