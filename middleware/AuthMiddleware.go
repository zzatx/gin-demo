package middleware

import (
	"cn.zzh.study/gin-demo/common"
	"cn.zzh.study/gin-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenStr := ctx.GetHeader("Authorization")

		//验证token格式
		if tokenStr == "" || strings.HasPrefix(tokenStr, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//提取token的有效部分
		tokenStr = tokenStr[7:]
		token, claims, err := common.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过，获取userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在
		ctx.Set("user", user)
		ctx.Next()
	}
}
