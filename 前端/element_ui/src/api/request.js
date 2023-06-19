import http from './http'
import qs from 'qs'
import axios from 'axios';
const ERR_OK = 0;
/** 
 * get方法，对应get请求 
 * @param {String} url [请求的url地址] 
 * @param {Object} params []
 */
//请求新闻数据
export function newsAPI(url) {
    return async function (num, channel) {
        return axios.get(url + '&num=' + num + '&channel=' + channel)
            .then(res => {
                return res.data.result.list;
            })
            .catch((e) => { return e });
    };
}

//version 2 另一种请求方法
/*
export function getarticle(params) {
   return request({
    method: 'GET',
    params,
   })
}
*/



/** 
 * post方法，对应post请求
 * @param {String} url [请求的url地址]
 * @param {Object} params [请求时携带的参数]
*/
//用户注册
export function userRegisterAPI(url) {
    return async function (form) {
        return axios.post(url, form)
            .then(res => {
                return res.data.msg;
            })
            .catch((e) => { return e });
    };
}
//用户登录
export function userLoginAPI(url) {
    return async function (form) {
        return axios.post(url, form)
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}
//修改信息
export function changeInfoAPI(url) {
    return async function (form, token) {
        const config = {
            headers: {
                Authorization: `Bearer ${token}`
            }
        };
        try {
            const res = await axios.get(url, form, config);
            return res.data;
        } catch (e) {
            return e;
        }
    };
}
//修改密码
export function changePasswordAPI(url) {
    return async function (form, token) {
        const config = {
            headers: {
                Authorization: `Bearer ${token}`
            }
        };
        try {
            const res = await axios.get(url, form, config);
            return res.data.msg;
        } catch (e) {
            return e;
        }
    };
}




//文章存储
export function storeNewsAPI(url) {
    return async function (item) {
        return axios.post(url, item, {
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}

//名字找文章
export function getNewsAPI(url) {
    return async function (title) {
        return axios.post(url, title
        )
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}