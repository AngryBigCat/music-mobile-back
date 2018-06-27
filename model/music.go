package model

type Music struct {
	Mid int `json:"id"`
	P_id int `json:"pid"`
	Name string `json:"name"`
	Num string `json:"number"`
	Site string `json:"filename"`
	Mytime string `json:"duration"`
	Att string `json:"remark"`
}

var (
	musics []Music
	music Music
)

func GetMusicList() []Music {
	db.Table("ml_detail").Find(&musics)
	return musics
}

func GetMusic(id string) Music {
	db.Table("ml_detail").Where("mid = ?", id).First(&music)
	return music
}

func GetMusicListIn(ids []int) []Music {
	db.Table("ml_detail").Where("mid in (?)", ids).Find(&musics)
	return musics
}
