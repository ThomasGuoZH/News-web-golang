package control

import (
	"backend/models"
	"backend/mysql"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 点赞处理
func LikeHandler(c *gin.Context) {
	var like *models.Likes
	if err := c.ShouldBindJSON(&like); err != nil {
		response.Fail(c, nil, "点赞失败")
		return
	}
	mysql.DB.Create(&like)
	var findComment models.Comment
	fmt.Println(like.CommentId)
	commentId, _ := strconv.Atoi(like.CommentId)
	result := mysql.DB.Where("ID=?", commentId).First(&findComment)
	if result.Error != nil {
		response.Fail(c, nil, "没有找到该评论")
	} else if result.Error == nil {
		fmt.Println(findComment)
		findComment.Likes++
		mysql.DB.Save(&findComment)
		response.Success(c, gin.H{
			"likes": findComment.Likes,
			"time":  findComment.Model.CreatedAt.String()[:19],
		}, "点赞成功")
	}
}
