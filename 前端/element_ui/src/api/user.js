import { userRegisterAPI, userLoginAPI } from "./request";

const userRegister = userRegisterAPI('http://127.0.0.1:8080/api/user/register')
const userLogin = userLoginAPI('http://127.0.0.1:8080/api/user/login')

export {
    userRegister,
    userLogin
}