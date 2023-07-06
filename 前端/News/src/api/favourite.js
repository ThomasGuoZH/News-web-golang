import { tokenAPI, getPersonalListAPI } from "./request";

//收藏
const favourite = tokenAPI('http://127.0.0.1:8080/api/user/create_fav')
//获取个人主页收藏列表
const getFavsList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/faves')
//判断是否收藏
const isFavs = tokenAPI('http://127.0.0.1:8080/api/user/isFaves')
//取消收藏
const disFave = tokenAPI('http://127.0.0.1:8080/api/user/dis_fav')
export {
    favourite,
    getFavsList,
    isFavs,
    disFave
}