package router

import (
	"backend/control"
	"backend/logic"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 用户路由组
func UserRouterInit(r *gin.RouterGroup) {
	userManager := r.Group("/user")
	{
		userManager.POST("/register", control.RegisterHandler)
		userManager.POST("/login", control.LoginHandler)
		userManager.Use(logic.AuthMiddleware())
		{
			userManager.POST("/change_info", control.ChangeInfoHandler)
			userManager.POST("/change_password", control.ChangePwdHandler)
		}
	}
}

// 新闻路由组
func NewsRouterInit(r *gin.RouterGroup) {
	newsManager := r.Group("/news")
	{
		newsManager.POST("/store_news", control.StoreHandler)
		newsManager.GET("/get_news", control.GetNewsHandler)
	}
}

// 评论路由组
func CommentRouterInit(r *gin.RouterGroup) {
	commentManager := r.Group("/comment")
	{
		//commentManager.GET("/comment_list", control.CommentListHandler)
		//userManager.Use(logic.AuthMiddleware())
		{
			commentManager.POST("/parent_comment", control.UserParentCommentHandler)
			commentManager.POST("/child_comment", control.UserChildCommentHandler)
			//commentManager.GET("/get_comment", control.GetCommentHandler)
		}
	}
}

//点赞路由组
/*func LikeRouterInit(r *gin.RouterGroup) {
	likeManager := r.Group("/like")
	{
		likeManager.Use(control.JWTAuthMiddleware){
			likeManager.POST("/like", control.LikeHandler)
			likeManager.GET("/get_like", control.GetLikeHandler)
			likeManager.GET("/like_count", control.LikeCountHandler)
	 	}
	}
}*/

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082", "http://localhost:8083",
		"http://localhost:8084", "http://localhost:8085"} // 允许访问的域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的HTTP方法
	// 配置 CORS 策略
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082", "http://localhost:8083",
		"http://localhost:8084", "http://localhost:8085"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))
	api := router.Group("api")
	UserRouterInit(api)
	NewsRouterInit(api)
	CommentRouterInit(api)
	return router
}
