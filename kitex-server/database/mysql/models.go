package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Video_ID int
	User_id  int
}

type Video struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	AuthorID uint `gorm:"foreignKey:AuthorID;references:User.User_id"` // 外键关联用户
	PlayURL  string
	Title    string
	CoverURL string

	FavoritedCount int64
	CommentCount   int64
}

type Comment struct {
	ID        uint `gorm:"primaryKey"` // 自增ID作为主键
	VideoID   uint `gorm:"foreignKey:VideoID;references:Video.ID"`
	UserID    uint `gorm:"foreignKey:UserID;references:User.ID"`
	Content   string
	CreatedAt time.Time // 评论时间
}

type Favorites struct {
	ID        int64          `json:"id"`
	UserID    int64          `json:"user_id"`
	VideoID   int64          `json:"video_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}
