package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../前端/html/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "首页.html", gin.H{})
	})
	router.Run(":80")
}
