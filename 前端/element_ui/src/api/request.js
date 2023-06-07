import http from './http'
/** 
 * get方法，对应get请求 
 * @param {String} url [请求的url地址] 
 * @param {Object} params [请求时携带的参数] 
 */
//请求新闻标题
export function getarticleTitleapi(url) {
    return function (params) {
        return new Promise((resolve, reject) => {
            axios.get(url + 'channel=' + params)
                .then(res => {
                    resolve(res.result.list.title);
                })
                .catch(err => {
                    reject(err.list)
                })
        });
    }

}
//请求新闻正文
export function getarticleContentapi(url) {
    return function (params) {
        return new Promise((resolve, reject) => {
            axios.get(url + 'channel=' + params)
                .then(res => {
                    resolve(res.result.list.content);
                })
                .catch(err => {
                    reject(err.list)
                })
        });
    }

}
//请求新闻照片地址
export function getarticlePicapi(url) {
    return function (params) {
        return new Promise((resolve, reject) => {
            axios.get(url + 'channel=' + params)
                .then(res => {
                    resolve(res.result.list.pic);
                })
                .catch(err => {
                    reject(err.list)
                })
        });
    }

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
