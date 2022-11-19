package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
}

func (t Product) TableName() string {
	return "product"
}

type UserInfo struct {
	ID    uint
	Name  string
	Age   uint
	Hobby string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/myDatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 判断是否存在
	if err = db.AutoMigrate(&UserInfo{}); err != nil {
		panic(err)
	}

	// 数据的插入
	u1 := UserInfo{1, "zyq", 18, "ow"}
	db.Create(&u1)
	u2 := UserInfo{2, "gyx", 18, "violin"}
	db.Create(&u2)

	// 用于数据查询的
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	// 根据主键进行删除
	db.Delete(&UserInfo{}, "2")

	// 对数据进行更新
	db.Model(&UserInfo{}).Where("id = ?", 1).Update("hobby", "apex")
}
