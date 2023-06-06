import axios from 'axios'




const request = axios.create({
    baseURL: 'https://api.jisuapi.com/news/get?appkey=14b3d45a212c10d0&start=0&num=20',
})

export default request