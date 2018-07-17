package model

type Music struct {
	Mid int `json:"id"`
	Name string `json:"name"`
	Num string `json:"number"`
	Site string `json:"filename"`
	Mytime string `json:"duration"`
	Att string `json:"remark"`
	Class string `json:"class"`
	Price float32 `json:"price"`
	Noprice float32 `json:"noprice"`
}

var (
	musics []Music
	music Music
)

func GetMusic(id string) Music {
	db.Table("ml_detail").
		Where("mid = ?", id).
		First(&music)
	return music
}

func GetMusicsOrder(order string, limit int) []Music {
	db.Table("ml_detail").
		Select(`
			ml_detail.*, 
			ml_class.class
		`).
		Joins("left join ml_class on ml_class.cid=ml_detail.p_id").
		Order(order).
		Limit(limit).
		Find(&musics)
	return musics
}

func GetMusicsLike(like string) []Music {
	db.Table("ml_detail").
		Select(`
			ml_detail.*, 
			ml_class.class
		`).
		Joins("left join ml_class on ml_class.cid=ml_detail.p_id").
		Where("name like ?", "%" + like +"%").
		Find(&musics)
	return musics
}

func GetMusicsIn(ids []string) []Music {
	db.Table("ml_detail").
		Where("mid in (?)", ids).
		Find(&musics)
	return musics
}