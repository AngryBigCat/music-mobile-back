package router

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/controller/v1"
)


func Init() *gin.Engine {
	r := gin.Default()

	v := r.Group("/api/v1")
	{
		v.GET("/music", v1.GetMusicList)
		v.GET("/music/:id", v1.GetMusic)
	}

	return r
}


