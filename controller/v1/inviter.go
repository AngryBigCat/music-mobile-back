package v1

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/model"
)

func GetInviters(c *gin.Context) {
	inviters := model.GetInviters()
	c.JSON(200, gin.H{
		"code": 200,
		"inviters": inviters,
	})
}