package main

import (
	"backend/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "root:123456@/practical-training?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
}
