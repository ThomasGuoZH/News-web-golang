package control

import (
	"backend/models"
	"backend/mysql"
	"backend/response"

	"github.com/gin-gonic/gin"
	//"github.com/go-sql-driver/mysql"
)

// 存储新闻文章
func StoreHandler(c *gin.Context) {
	var news *models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		response.Fail(c, nil, "新闻存储失败")
		return
	}
	var findNews models.News
	result := mysql.DB.Where("title=?", news.Title).First(&findNews)
	if result.Error == nil {
		response.Fail(c, nil, "新闻已存在")
		return
	}
	mysql.DB.Create(&news)
	response.Success(c, nil, "新闻存储成功")

}

// 获取文章
func GetNewsHandler(c *gin.Context) {
	var getNews *models.GetNews
	if err := c.ShouldBindJSON(&getNews); err != nil {
		response.Fail(c, nil, "新闻获取失败，未提供正确的Title")
		return
	}
	var findNews models.News
	result := mysql.DB.Where("title=?", getNews.Title).First(&findNews)
	if result.Error != nil {
		response.Fail(c, nil, "新闻不存在!")
		return
	}
	newsData := gin.H{
		"title":   findNews.Title,
		"content": findNews.Content,
		"pic":     findNews.Pic,
		"time":    findNews.Time,
		"src":     findNews.Src,
	}
	response.Success(c, newsData, "新闻获取成功")
}
