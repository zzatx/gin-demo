package main

import (
	"cn.zzh.study/gin-demo/common"
	"cn.zzh.study/gin-demo/model"
	"github.com/gin-gonic/gin"
)

func main() {
	var objects = []interface{}{&model.User{}}
	common.InitDB(objects)
	r := gin.Default()
	CollectRoute(r)
	panic(r.Run(":8089"))
}
