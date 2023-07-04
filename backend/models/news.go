package models

type News struct {
	Title    string `json:"title" db:"title" binding:"required"`
	Content  string `json:"content" db:"content" gorm:"type:text" binding:"required"`
	Pic      string `json:"pic" db:"pic" binding:"required"`
	Time     string `json:"time" db:"time" binding:"required"`
	Src      string `json:"src" db:"src" binding:"required"`
	Category string `json:"category" db:"category"`
	Url      string `json:"url" db:"url"`
	Weburl   string `json:"weburl" db:"weburl"`
}

// 数据表名字
func (News) TableName() string {
	return "news"
}
