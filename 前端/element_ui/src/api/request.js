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
export function articleAPI(url) {
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



/*
 * post方法，对应post请求
 * @param {String} url [请求的url地址]
 * @param {Object} params [请求时携带的参数]
 */
// export function post(url, params) {
//     return new Promise((resolve, reject) => {
//         axios.post(url, QS.stringify(params))
//             .then(res => {
//                 resolve(res.data);
//             })
//             .catch(err => {
//                 reject(err.data)
//             })
//     });
// }
