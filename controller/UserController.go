package controller

import (
	"cn.zzh.study/gin-demo/common"
	"cn.zzh.study/gin-demo/model"
	"cn.zzh.study/gin-demo/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

var DB = common.GetDB()

func Register(ctx *gin.Context) {
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
		name = utils.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"msg": "用户已存在"})
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密错误"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword)}
	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{"code": 200, "msg": "注册成功"})
}

func Login(ctx *gin.Context) {
	//获取参数
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
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).Find(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"msg": "密码错误请确认后重试"})
		return
	}
	//TODO 发放token
	token := "111"

	//返回结果
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登陆成功"})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).Find(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
