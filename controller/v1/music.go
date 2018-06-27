package v1

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/model"
)

func GetMusicList(c *gin.Context) {
	action := c.Query("action")
	var musics []model.Music
	switch {
	case action == "hot":
		musics = model.GetMusicListIn([]int{36})
	case action == "latest":
		musics = model.GetMusicListIn([]int{37})
	}
	c.JSON(200, gin.H{
		"code": 0,
		"musics": musics,
	})
}

func GetMusic(c *gin.Context) {
	id := c.Param("id")
	music := model.GetMusic(id)
	c.JSON(200, gin.H{
		"code": 0,
		"music": music,
	})
}
