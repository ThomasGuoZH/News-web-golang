package models

type News struct {
	Title   string `json:"title" db:"title" binding:"required"`
	Content string `json:"content" db:"content" binding:"required"`
	Channel string `json:"channel" db:"channel"`
	Pic     string `json:"pic" db:"pic"`
	Time    string `json:"time" db:"time"`
	Src     string `json:"src" db:"src"`
}

// 数据表名字
func (News) TableName() string {
	return "news"
}

type GetNews struct {
	Title string `json:"title" binding:"required"`
}
