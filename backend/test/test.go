package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../前端/html/*")
	router.GET("", func(context *gin.Context) {
		context.HTML(http.StatusOK, "首页.html", gin.H{})
	})
	router.GET("/体育.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "体育.html", gin.H{})
	})
	router.Run(":80")
}
