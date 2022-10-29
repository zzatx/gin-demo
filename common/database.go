package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(objects []interface{}) *gorm.DB {
	host := "169.254.142.159"
	port := "3306"
	database := "go-study"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		//有错误连接失败则直接终止程序
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(objects...)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
