const { defineConfig } = require('@vue/cli-service')
module.exports = {
  devServer: {
    open: true,
    host: 'localhost',
    port: 8080,
    https: false,
    //以上的ip和端口是我们本机的;下面为需要跨域的
    proxy: { //配置跨域
      '/api': {
        target: 'https://api.jisuapi.com/news/get?appkey=cc5e3b31d2129d83&start=0', //填写请求的目标地址
        changOrigin: true, //允许跨域
        pathRewrite: {
          '^/api': 'newsapi' //请求的时候使用这个api就可以
        }
      }
    }
  }
}
// cc5e3b31d2129d83
// 14b3d45a212c10d0
// 52c8aff9fab2d048
// ec0ed7787af13ce3
// faec4ccb9f663492
// 36f96971fb8e5a8e
// 37c8518088eabae4
// 71f1655f9b3b89e4

//37c8518088eabae4
//d03f43c234cd72fb