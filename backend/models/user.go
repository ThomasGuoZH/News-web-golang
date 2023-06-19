package models

// 用户基本属性
type User struct {
	Id       uint64 `json:"id,string" db:"id"`
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

// RegisterForm 注册请求参数
type RegisterForm struct {
	UserName        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
	Sex             string `form:"sex"`
	Email           string `form:"email"`
	Phone           string `form:"phone"`
}

// LoginForm 登录请求参数
type LoginForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

// ChangeInfoForm 更改信息请求参数(需要Id)
type ChangeInfoForm struct {
	Id       string `form:"id"`
	UserName string `form:"username"`
	Sex      string `form:"sex"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
}

// ChangePwdForm 更改密码请求参数(需要Id)
type ChangePwdForm struct {
	Id              string `form:"id"`
	OldPassword     string `form:"oldPassword"`
	NewPassword     string `form:"newPassword"`
	ConfirmPassword string `form:"confirmPassword"`
}
