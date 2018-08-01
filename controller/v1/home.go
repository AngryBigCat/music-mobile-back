package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/AngryBigCat/music-mobile-back/model"
	"net/http"
)

func GetHomeText(c *gin.Context) {
	text := model.GetHomeText()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"text": text.Text,
	})
}