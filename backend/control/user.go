package control

import (
	"backend/logic"
	"backend/models"
	"backend/mysql"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/go-sql-driver/mysql"
)

// 注册
func RegisterHandler(c *gin.Context) {
	var userRegister *models.RegisterForm
	if err := c.ShouldBind(&userRegister); err != nil {
		response.Fail(c, nil, "注册失败")
	} else {
		//判断用户名是否存在
		var findUser models.User
		result := mysql.DB.Where("user_name=?", userRegister.UserName).First(&findUser)
		if result.Error == nil {
			response.Fail(c, nil, "用户名已被使用!")
		} else if result.Error != nil {
			// 生成随机唯一id
			var randomId = logic.GenerateID()
			fmt.Println(randomId)
			var newUser = models.User{
				Id:       randomId,
				UserName: userRegister.UserName,
				Password: userRegister.Password,
				Phone:    userRegister.Phone,
				Sex:      userRegister.Sex,
				Email:    userRegister.Email,
			}
			mysql.DB.Create(&newUser)
			response.Success(c, nil, "注册成功!")
		}
	}
}

// 登录
func LoginHandler(c *gin.Context) {
	var userLogin *models.LoginForm
	if err := c.ShouldBind(&userLogin); err != nil {
		response.Fail(c, nil, "登录失败")
	} else {
		var findUser models.User
		result := mysql.DB.Where("user_name=?", userLogin.UserName).First(&findUser)
		if result.Error != nil {
			response.Fail(c, nil, "用户不存在!")
		} else if result.Error == nil {
			// 查找密码
			if findUser.Password != userLogin.Password {
				response.Fail(c, nil, "密码错误!")
			} else {
				token, err := logic.ReleaseToken(findUser) // 发放token
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "System err!"})
					fmt.Printf("token generate error:%v", err)
					return
				}
				response.Success(c, gin.H{
					"token":    token,
					"userId":   fmt.Sprintf("%d", findUser.Id),
					"username": findUser.UserName,
					"sex":      findUser.Sex,
					"email":    findUser.Email,
					"phone":    findUser.Phone,
				}, "登录成功!")
			}
		}
	}
}

// 获得用户信息
func UserInfoHandler() {

}
