package control

import (
	"backend/models"
	"backend/mysql"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// 获取评论列表
func CommentListHandler() {

}

// 用户1级评论
func UserParentCommentHandler(c *gin.Context) {
	var parentComment models.ParentComment
	//format error
	if err := c.ShouldBindJSON(&parentComment); err != nil {
		response.Fail(c, nil, "评论失败")
		return
	}
	var comment models.Comment
	var randID = uint(rand.Uint32())
	comment.Model.ID = randID
	comment.ParentId = randID
	comment.Type = 0
	comment.Likes = 0
	comment.Title = parentComment.Title
	comment.Content = parentComment.Content
	comment.Author = parentComment.Author
	comment.Channel = parentComment.Channel
	mysql.DB.Create(&comment)
	fmt.Println(comment)
	response.Success(c, gin.H{
		"author":  comment.Author,
		"content": comment.Content,
		"id":      strconv.FormatInt(int64(comment.Model.ID), 10),
		"likes":   comment.Likes,
		"time":    comment.Model.CreatedAt.String()[:19],
	}, "评论成功")
}

// 用户2级评论
func UserChildCommentHandler(c *gin.Context) {
	var comment *models.ChildComment
	//format error
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Fail(c, nil, "评论失败")
		return
	}
	parentId, errId := strconv.ParseUint(comment.ParentId, 10, 32)
	if errId != nil {
		response.Fail(c, nil, "1级评论不存在")
		return
	}
	var findParent models.Comment
	result := mysql.DB.Where("parent_id=?", parentId).First(&findParent)
	if result.Error != nil {
		response.Fail(c, nil, "1级评论不存在")
		return
	}
	var newChild models.Comment
	newChild.Model.ID = uint(rand.Uint32())
	newChild.Title = comment.Title
	newChild.Author = comment.Author
	newChild.Content = comment.Content
	newChild.Channel = comment.Channel
	newChild.Type = 1
	newChild.ParentId = uint(parentId)
	newChild.Likes = 0
	mysql.DB.Create(&newChild)
	response.Success(c, gin.H{
		"author":  newChild.Author,
		"content": newChild.Content,
		"id":      strconv.FormatInt(int64(newChild.Model.ID), 10),
		"likes":   newChild.Likes,
		"time":    newChild.Model.CreatedAt.String()[:19],
	}, "评论成功")
}

// 点赞处理
func LikeHandler(c *gin.Context) {
	var like *models.Likes
	if err := c.ShouldBindJSON(&like); err != nil {
		response.Fail(c, nil, "点赞失败")
		return
	}
	var findComment models.Comment
	fmt.Println(like.CommentId)
	commentId, _ := strconv.Atoi(like.CommentId)
	result := mysql.DB.Where("ID=?", commentId).First(&findComment)
	if result.Error != nil {
		response.Fail(c, nil, "没有找到该评论")
	} else if result.Error == nil {
		var likedByUser models.Likes
		result = mysql.DB.Where("comment_id = ? AND liker = ?", commentId, like.Liker).First(&likedByUser)
		if result.RowsAffected > 0 {
			response.Fail(c, nil, "你已经点赞过该评论")
			return
		}
		mysql.DB.Create(&like)
		fmt.Println(findComment)
		findComment.Likes++
		mysql.DB.Save(&findComment)
		response.Success(c, gin.H{
			"likes": findComment.Likes,
			"time":  findComment.Model.CreatedAt.String()[:19],
		}, "点赞成功")
	}
}
