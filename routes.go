package main

import (
	"cn.zzh.study/gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/suth/login", controller.Login)
	return r
}
