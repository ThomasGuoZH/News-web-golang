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
        const headers = {
            'Authorization': `Bearer ${token}`
        }
        try {
            const res = await axios.post(url, form, { headers });
            return res.data;
        } catch (e) {
            return e;
        }
    };
}
//修改密码
export function changePasswordAPI(url) {
    return async function (form, token) {
        const headers = {
            'Authorization': `Bearer ${token}`
        }
        try {
            const res = await axios.post(url, form, { headers });
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
                return res.data.msg;
            })
            .catch((e) => { return e });
    };
}

//名字找文章
export function getNewsAPI(url) {
    return async function (title) {
        return axios.get(url, {
            params: { title: title }
        })
            .then(res => {
                return res.data.data;
            })
            .catch((e) => { return e });
    };
}

//发布父评论
export function parentCommentAPI(url) {
    return async function (parentComment, token) {
        const headers = {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
        return axios.post(url, parentComment, { headers })
            .then(res => {
                return res.data;
            }).catch((e) => { return e });
    };
}

//发布子评论
export function childCommentAPI(url) {
    return async function (reply, token) {
        const headers = {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
        return axios.post(url, reply, { headers })
            .then(res => {
                return res.data;
            }).catch((e) => { return e });
    };
}

// 获取新闻评论列表
export function getNewsCommentsListAPI(url) {
    return async function (title) {
        return axios.get(url, {
            params: { title: title }
        })
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}

//点赞
export function LikesAPI(url) {
    return async function (like, token) {
        const headers = {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
        return axios.post(url, like, { headers })
            .then(res => {
                return res.data;
            }).catch((e) => { return e });
    };
}


//收藏新闻
export function favouriteAPI(url) {
    return async function (favourite, token) {
        const headers = {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        }
        return axios.post(url, favourite, { headers })
            .then(res => {
                return res.data;
            }).catch((e) => { return e });
    };
}


//获得个人评论列表
export function getPersonalCommentsListAPI(url) {
    return async function (id) {
        return axios.get(url, {
            params: { user_id: id }
        })
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}

//获取个人点赞列表
export function getLikesListAPI(url) {
    return async function (id) {
        return axios.get(url, {
            params: { user_id: id }
        })
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}

//获得个人回复列表
export function getRepliesListAPI(url) {
    return async function (id) {
        return axios.get(url, {
            params: { user_id: id }
        })
            .then(res => {
                return res.data;
            })
            .catch((e) => { return e });
    };
}