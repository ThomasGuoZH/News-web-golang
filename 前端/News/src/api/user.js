import { userRegisterAPI, userLoginAPI, changeInfoAPI, changePasswordAPI } from "./request";

//注册
const userRegister = userRegisterAPI('http://127.0.0.1:8080/api/user/register')
//登录
const userLogin = userLoginAPI('http://127.0.0.1:8080/api/user/login')
//修改信息
const changeInfo = changeInfoAPI('http://127.0.0.1:8080/api/user/change_info')
//修改密码
const changePassword = changePasswordAPI('http://127.0.0.1:8080/api/user/change_password')

export {
    userRegister,
    userLogin,
    changeInfo,
    changePassword
}