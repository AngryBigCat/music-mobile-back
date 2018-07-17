package v1

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/model"
)

func GetOrder(c *gin.Context) {
	token := c.Query("token")
	user, _ := model.ParseToken(token)
	print(user.Uid)
}