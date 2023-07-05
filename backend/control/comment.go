package control

import (
	"backend/logic"
	"backend/models"
	"backend/mysql"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewsExists(title string) bool {
	var news models.News
	result := mysql.DB.Where("title=?", title).First(&news)
	if result.Error != nil {
		return false
	}
	return true
}

func UserExists(name string) bool {
	var user models.User
	result := mysql.DB.Where("user_name=?", name).First(&user)
	if result.Error != nil {
		return false
	}
	return true
}

// 获取新闻评论列表
func CommentListHandler(c *gin.Context) {
	title := c.Query("title")
	if !NewsExists(title) {
		response.Fail(c, nil, "新闻不存在")
		return
	}
	var comments, subComments []models.Comment
	// find corresponding comment
	mysql.DB.Find(&comments, "type=0 AND title=?", title)
	var commentsArr, subCommentsArr []gin.H
	for _, comment := range comments {
		// find corresponding sub comment
		mysql.DB.Find(&subComments, "type=1 AND parent_id=?", comment.ID)
		for _, subComment := range subComments {
			// append corresponding sub comment
			subCommentsArr = append(subCommentsArr, gin.H{
				"author":  subComment.Author,
				"content": subComment.Content,
				"id":      strconv.FormatInt(int64(subComment.ID), 10),
				"likes":   subComment.Likes,
				"time":    subComment.CreatedAt.String()[:19],
			})
		}
		// append comment
		commentsArr = append(commentsArr, gin.H{
			"author":  comment.Author,
			"content": comment.Content,
			"id":      strconv.FormatInt(int64(comment.ID), 10),
			"likes":   comment.Likes,
			"time":    comment.CreatedAt.String()[:19],
			"replies": subCommentsArr,
		})
		// clear subComments, subCommentsArr
		subComments = subComments[0:0]
		subCommentsArr = nil
	}
	response.Success(c, gin.H{
		"comments": commentsArr,
	}, "获取所有评论成功")
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
	var randID = logic.GenerateID32()
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
	newChild.Model.ID = logic.GenerateID32()
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

// 用户ID对应1级评论列表
func PersonalCommentsListHandler(c *gin.Context) {
	idString := c.Query("user_id")
	if len(idString) == 0 {
		response.Fail(c, nil, "用户不存在")
		return
	}
	var author string
	if id, err := strconv.ParseUint(idString, 10, 64); err != nil {
		fmt.Println(1)
		response.Fail(c, nil, "用户不存在")
		return
	} else {
		var user models.User
		result := mysql.DB.Where("id=?", id).First(&user)
		if result.Error != nil {
			fmt.Println(1)
			response.Fail(c, nil, "用户不存在")
			return
		}
		author = user.UserName
	}
	var comments []models.Comment
	// find comments with corresponding author
	mysql.DB.Find(&comments, "type=0 AND author=?", author)
	var commentsArr []gin.H
	for _, comment := range comments {
		commentsArr = append(commentsArr, gin.H{
			"content": comment.Content,
			"time":    comment.CreatedAt.String()[:19],
			"channel": comment.Channel,
			"title":   comment.Title,
		})
	}
	response.Success(c, gin.H{"comments": commentsArr}, "获取用户评论成功")
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

// 收到的赞
func PersonalLikesListHandler(c *gin.Context) {
	idString := c.Query("user_id")
	if len(idString) == 0 {
		response.Fail(c, nil, "用户不存在")
		return
	}
	var author string
	if id, err := strconv.ParseUint(idString, 10, 64); err != nil {
		fmt.Println(1)
		response.Fail(c, nil, "用户不存在")
		return
	} else {
		var user models.User
		result := mysql.DB.Where("id=?", id).First(&user)
		if result.Error != nil {
			fmt.Println(1)
			response.Fail(c, nil, "用户不存在")
			return
		}
		author = user.UserName
	}
	var likes []models.Likes
	// find comments with corresponding author
	mysql.DB.Find(&likes, "author=?", author)
	var likesArr []gin.H
	for _, like := range likes {
		likesArr = append(likesArr, gin.H{
			"liker":   like.Liker,
			"content": like.Content,
			"time":    like.CreatedAt.String()[:19],
			"channel": like.Channel,
			"title":   like.Title,
		})
	}
	response.Success(c, gin.H{"likes": likesArr}, "获取用户评论成功")
}
