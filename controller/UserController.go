package controller

import (
	"cn.zzh.study/gin-demo/common"
	"cn.zzh.study/gin-demo/model"
	"cn.zzh.study/gin-demo/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password}
	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{"msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).Find(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
