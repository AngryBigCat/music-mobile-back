package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var r *gin.Engine

func Init() *gin.Engine {
	// 初始化路由
	r = gin.Default()

	// 跨域设置
	r.Use(cors.Default())

	// 注册 v1 路由组
	v1Group()

	return r
}


