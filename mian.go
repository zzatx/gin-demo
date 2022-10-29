package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	DB := InitDB(&User{})
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "手机号必须为11位"})
			return
		}
		if len(password) <= 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "密码不能小于6位"})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		//判断手机号是否存在
		if isTelephoneExist(DB, telephone) {
			ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"msg": "用户已存在"})
			return
		}
		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password}
		DB.Create(&newUser)
		//返回结果
		ctx.JSON(200, gin.H{"msg": "注册成功"})
	})
	panic(r.Run(":8089"))
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).Find(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	//不是用seed则每次生成的随机数都将一样
	rand.Seed(time.Now().Unix())
	for i := range result {
		//rand.Intn(n) 随机生成一个0到n-1的随机数 0 <= i < n
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB(object interface{}) *gorm.DB {
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
	db.AutoMigrate(object)
	return db

}
