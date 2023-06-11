package models

type Article struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Channel string `json:"channel" db:"channel"`
	Pic     string `json:"pic" db:"pic"`
	Time    string `json:"time" db:"time"`
	Src     string `json:"src" db:"src"`
}
