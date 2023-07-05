import { tokenAPI, getPersonalListAPI } from "./request";

const favourite = tokenAPI('http://127.0.0.1:8080/api/user/create_fav')
const getFavsList = getPersonalListAPI('http://127.0.0.1:8080/api/personal/faves')
const isFavs = tokenAPI('http://127.0.0.1:8080/api/user/isFaves')

export {
    favourite,
    getFavsList,
    isFavs
}