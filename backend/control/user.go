package control

import (
	"backend/logic"
	"backend/models"
	"backend/mysql"
	"backend/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func ChangeInfoHandler(c *gin.Context) {
	var newInfo models.ChangeInfoForm
	if err := c.ShouldBind(&newInfo); err != nil {
		response.Fail(c, nil, "更改用户信息失败")
		return
	}
	var findUser, updateUser, findName models.User
	Id, _ := strconv.ParseUint(newInfo.Id, 10, 64)
	result := mysql.DB.Where("id=?", Id).First(&findUser)
	if result.Error != nil {
		response.Fail(c, nil, "用户不存在!")
		return
	}
	result = mysql.DB.Where("user_name=?", newInfo.UserName).First(&findName)
	// 用户名重复
	if findName.UserName == newInfo.UserName && findName.Id != Id {
		response.Fail(c, nil, "用户名已存在")
		return
	}
	updateUser.Id = findUser.Id
	updateUser.Password = findUser.Password
	updateUser.UserName = newInfo.UserName
	updateUser.Sex = newInfo.Sex
	updateUser.Email = newInfo.Email
	updateUser.Phone = newInfo.Phone
	if updateUser == findUser {
		response.Fail(c, nil, "用户信息不能相同")
		return
	}
	mysql.DB.Save(&updateUser)
	response.Success(c, gin.H{
		"username": newInfo.UserName,
		"sex":      newInfo.Sex,
		"email":    newInfo.Email,
		"phone":    newInfo.Phone,
	}, "保存成功!")
}

func ChangePwdHandler(c *gin.Context) {
	var newPwd models.ChangePwdForm
	if err := c.ShouldBind(&newPwd); err != nil {
		response.Fail(c, nil, "更改密码失败")
		return
	}
	var findUser models.User
	Id, _ := strconv.ParseUint(newPwd.Id, 10, 64)
	result := mysql.DB.Where("id=?", Id).First(&findUser)
	if result.Error != nil {
		response.Fail(c, nil, "用户不存在!")
		return
	}
	if findUser.Password != newPwd.OldPassword {
		response.Fail(c, nil, "旧密码错误")
		return
	}
	if findUser.Password == newPwd.NewPassword {
		response.Fail(c, nil, "新旧密码不能相同")
		return
	}
	findUser.Password = newPwd.NewPassword
	mysql.DB.Save(&findUser)
	response.Success(c, gin.H{
		"username": findUser.UserName,
		"sex":      findUser.Sex,
		"email":    findUser.Email,
		"phone":    findUser.Phone}, "密码修改成功!")
}
