package main

import (
	"cn.zzh.study/gin-demo/controller"
	"cn.zzh.study/gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/suth/login", controller.Login)
	//middleware.AuthMiddleware()为中间件，类似于java中的拦截器
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
