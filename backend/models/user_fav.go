package models

import "time"

func (UserFav) TableName() string {
	return "user_fav"
}

type UserFav struct {
	Title     string `db:"title" gorm:"primaryKey"`
	UserId    uint64 `db:"user_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	Channel   string `db:"channel"`
}

// NewUserFav 需要channel title user_id
type NewUserFav struct {
	Channel string `json:"channel" binding:"required"`
	Title   string `json:"title" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}
