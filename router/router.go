package router

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/controller/v1"
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

/*
  定义api v1路由组
  需要将该函数注册在Init()方法里
 */
func v1Group() {
	v := r.Group("/api/v1")
	{
		v.GET("/music", v1.GetMusicList)
		v.GET("/music/:id", v1.GetMusic)
	}
}


