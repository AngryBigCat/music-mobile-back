package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/AngryBigCat/music-mobile-back/model"
	"strings"
)


type HomeMusic struct {
	Name string `json:"name"`
	Musics []model.Music `json:"musics"`
}


func GetMusic(c *gin.Context) {
	id := c.Param("id")
	music := model.GetMusic(id)
	c.JSON(200, gin.H{
		"code": 0,
		"music": music,
	})
}

/**
QueryString
ids=1,2,3,4,5
search=金蛇狂舞
page=home
 */
func GetMusics(c *gin.Context) {
	var musics []model.Music
	ids := c.Query("ids")
	search := c.Query("search")
	page := c.Query("page")
	switch {
		case ids != "":
			idsArr := strings.Split(ids, ",")
			musics = model.GetMusicsIn(idsArr)
		case search != "":
			musics = model.GetMusicsLike(search)
		case page != "":
			if page == "home" {
				musics := getHomeMusics()
				c.JSON(200, gin.H{
					"code": 0,
					"musics": musics,
				})
				return
			}
	}

	c.JSON(200, gin.H{
		"code": 0,
		"musics": musics,
	})
}


func getHomeMusics() []HomeMusic {
	musics := []HomeMusic{
		HomeMusic{"热门曲目", model.GetMusicsOrder("mid asc", 10)},
		HomeMusic{"最新曲目", model.GetMusicsOrder("time desc", 10)},
	}
	return musics
}
