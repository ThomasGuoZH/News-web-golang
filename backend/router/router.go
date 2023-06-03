package main

import (
	"github.com/gin-gonic/gin"
)

// 用户路由组
func UserRouterInit(r *gin.RouterGroup) {
	/*userManager := r.Group("/users")
	{
		userManager.GET("/register", RegisterHandler())
		userManager.GET("/login", LoginHandler())

	}*/
}

// 新闻路由组
func ArticleRouterInit(r *gin.RouterGroup) {
	articleManager := r.Group("/users")
	{
		articleManager.GET("/register")
		articleManager.GET("/login")
		//
	}
}

func main() {
	router := gin.Default()
	api := router.Group("api")
	UserRouterInit(api)
	ArticleRouterInit(api)

	err := router.Run(":80")
	if err != nil {
		return
	}
}
