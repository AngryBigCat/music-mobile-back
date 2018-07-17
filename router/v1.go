package router

import (
	"music-mobile-back/controller/v1"
)

/*
  定义api v1路由组
  需要将该函数注册在Init()方法里
 */
func v1Group() {
	v := r.Group("/api/v1")

	v.POST("/token", v1.GetToken)
	v.POST("/register", v1.Register)
	v.GET("/getCode", v1.GetCode)

	v.GET("/musics", v1.GetMusics)
	v.GET("/music/:id", v1.GetMusic)

	v.GET("/inviters", v1.GetInviters)

	v.GET("/order", v1.GetOrder)
}