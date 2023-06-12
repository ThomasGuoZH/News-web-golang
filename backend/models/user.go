package models

import (
	"github.com/jinzhu/gorm"
)

// 用户基本属性
type User struct {
	gorm.Model
	Id       string `json:"id,string" db:"id"`
	UserName string `json:"username" db:"user_name"`
	Password string `json:"password" db:"password"`
	Sex      string `json:"sex" db:"sex"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
}

// 数据表名字
func (User) TableName() string {
	return "user"
}

// 用户构造函数
func NewUser(name, password, sex, email, phone string) *User {
	return &User{
		UserName: name,
		Password: password,
		Sex:      sex,
		Email:    email,
		Phone:    phone,
	}
	//生成id。
}

// RegisterForm 登录请求参数
type RegisterForm struct {
	UserName        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
	Sex             string `form:"sex"`
	Email           string `form:"email"`
	Phone           string `form:"phone"`
}

// 注册构造函数
func NewRegisterForm(name, password, sex, email, phone string) *RegisterForm {
	return &RegisterForm{
		UserName: name,
		Password: password,
		Sex:      sex,
		Email:    email,
		Phone:    phone,
	}
}

// LoginForm 登录请求参数
type LoginForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}
