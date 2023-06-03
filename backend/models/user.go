package models

import (
	"github.com/jinzhu/gorm"
)

// 用户基本属性
type User struct {
	gorm.Model
	Name     string `json:"用户名" db:"name"`
	Id       string `json:"账号" db:"id"`
	Password string `json:"密码" db:"password"`
	Level    int    `json:"等级" db:"level"`
	Sex      string `json:"性别" db:"sex"`
	Email    string `json:"邮箱" db:"email"`
	Phone    string `json:"电话号码" db:"phone"`
	Age      int    `json:"年龄" db:"age"`
}

// Register
type RegisterForm struct {
	Id              string `json:"账号" binding:"required"`
	Password        string `json:"密码" binding:"required"`
	ConfirmPassword string `json:"确认密码" binding:"eqfiled=Password" binding:"required"`
}
