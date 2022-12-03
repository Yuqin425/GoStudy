package dao

import (
	"goPractice/Week_08/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:zyq425ZYQ@(127.0.0.1:3306)/myDatabase?charset=utf8&parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	//迁移
	db.AutoMigrate(&model.User{})

	DB = db

	return db
}
