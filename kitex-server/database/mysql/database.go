package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:tcx661053495996@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{}) // 自动迁移
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&Comment{})
	return db
}
