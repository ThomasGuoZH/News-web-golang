import { userRegisterAPI, userLoginAPI, changeInfoAPI, changePasswordAPI } from "./request";

const userRegister = userRegisterAPI('http://127.0.0.1:8080/api/user/register')
const userLogin = userLoginAPI('http://127.0.0.1:8080/api/user/login')
const changeInfo = changeInfoAPI('http://127.0.0.1:8080/api/user/change_info')
const changePassword = changePasswordAPI('http://127.0.0.1:8080/api/user/change_info')

export {
    userRegister,
    userLogin,
    changeInfo,
    changePassword
}