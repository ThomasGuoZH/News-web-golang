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
	} else {
		var newNews = models.News{
			Title:   news.Title,
			Content: news.Content,
			Channel: news.Channel,
			Pic:     news.Pic,
			Time:    news.Time,
			Src:     news.Src,
		}
		mysql.DB.Create(&newNews)
		response.Success(c, nil, "新闻存储成功")
	}
}

// 获取文章
func GetNewsHandler(c *gin.Context) {
	var getNews *models.GetNews
	if err := c.ShouldBindJSON(&getNews); err != nil {
		response.Fail(c, nil, "新闻获取失败，未提供正确的Title")
	} else {
		var findNews models.News
		result := mysql.DB.Where("title=?", getNews.Title).First(&findNews)
		if result.Error != nil {
			response.Fail(c, nil, "新闻不存在!")
		} else {
			newsData := gin.H{
				"title":   findNews.Title,
				"content": findNews.Content,
				"pic":     findNews.Pic,
				"time":    findNews.Time,
				"src":     findNews.Src,
			}
			response.Success(c, newsData, "新闻获取成功")
		}

	}

}
