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
	var newsList []*models.News
	if err := c.ShouldBindJSON(&newsList); err != nil {
		response.Fail(c, nil, "新闻存储失败")
	} else {
		for _, news := range newsList {
			var count int
			if err := mysql.DB.Model(&models.News{}).Where("title = ?", news.Title).Count(&count).Error; err != nil {
				response.Fail(c, nil, "查找新闻失败")
			} else if count > 0 {
				continue
			} else if count == 0 {
				mysql.DB.Create(&news)
			}
		}
		response.Success(c, nil, "新闻存储成功")
	}
}

// 获取文章
func GetNewsHandler(c *gin.Context) {
	title := c.Query("title")
	var findNews models.News
	result := mysql.DB.Where("title=?", title).First(&findNews)
	if result.Error != nil {
		response.Fail(c, nil, "新闻不存在!")
	} else {
		response.Success(c, gin.H{
			"title":   findNews.Title,
			"content": findNews.Content,
			"pic":     findNews.Pic,
			"time":    findNews.Time,
			"src":     findNews.Src,
		}, "新闻获取成功")
	}
}
