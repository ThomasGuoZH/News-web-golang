package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func Init(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, hostname, port, dbname))
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

//func main() {
//	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
//	db, err := gorm.Open("mysql", "root:123456@/practical-training?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		panic(err)
//	}
//	db.AutoMigrate(&models.User{})
//	defer func(db *gorm.DB) {
//		err := db.Close()
//		if err != nil {
//
//		}
//	}(db)
//}
