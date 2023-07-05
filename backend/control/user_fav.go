package control

import (
	"backend/models"
	"backend/mysql"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserIdExists(id uint64) bool {
	var user models.User
	result := mysql.DB.Where("id=?", id).First(&user)
	if result.Error != nil {
		return false
	}
	return true
}

// 需要channel title user_id
func CreateFavHandler(c *gin.Context) {
	var fav *models.NewUserFav
	if err := c.ShouldBindJSON(&fav); err != nil {
		response.Fail(c, nil, "收藏失败")
		return
	}
	if !NewsExists(fav.Title) {
		response.Fail(c, nil, "新闻不存在")
		return
	}
	id, err := strconv.ParseUint(fav.UserId, 10, 64)
	if err != nil || !UserIdExists(id) {
		response.Fail(c, gin.H{
			"isFavorite": false,
		}, "用户不存在")
		return
	}
	var findFav models.UserFav
	result := mysql.DB.Where("title=? AND user_id=?", fav.Title, id).First(&findFav)
	if result.Error == nil {
		response.Fail(c, gin.H{
			"isFavorite": true,
		}, "已收藏过")
		return
	}
	createFav := models.UserFav{
		Title:   fav.Title,
		UserId:  id,
		Channel: fav.Channel,
	}
	mysql.DB.Create(&createFav)
	response.Success(c, gin.H{
		"isFavorite": true,
	}, "收藏成功")
}

// GetFavHandler 获取user_id收藏的news
func GetFavHandler(c *gin.Context) {
	idString := c.Query("user_id")
	if len(idString) == 0 {
		response.Fail(c, nil, "用户不存在")
		return
	}
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil || !UserIdExists(id) {
		response.Fail(c, nil, "用户不存在")
		return
	}
	var faves []models.UserFav
	var favesArr []gin.H
	mysql.DB.Find(&faves, "user_id=?", id)
	for _, val := range faves {
		favesArr = append(favesArr, gin.H{
			"title":   val.Title,
			"channel": val.Channel,
			"time":    val.CreatedAt.String()[:19],
		})
	}
	fmt.Println(favesArr)
	response.Success(c, gin.H{"favourites": favesArr}, "获取收藏成功")
}

// 判断是否收藏
func IsFavHandler(c *gin.Context) {
	var fav *models.NewUserFav
	if err := c.ShouldBindJSON(&fav); err != nil {
		response.Fail(c, nil, "绑定失败")
		return
	}
	if !NewsExists(fav.Title) {
		response.Fail(c, nil, "新闻不存在")
		return
	}
	id, err := strconv.ParseUint(fav.UserId, 10, 64)
	if err != nil || !UserIdExists(id) {
		response.Fail(c, gin.H{
			"isFavorite": false,
		}, "未收藏")
		return
	}
	var findFav models.UserFav
	result := mysql.DB.Where("title=? AND user_id=?", fav.Title, id).First(&findFav)
	if result.Error == nil {
		response.Success(c, gin.H{
			"isFavorite": true,
		}, "已收藏")
		return
	}
}
