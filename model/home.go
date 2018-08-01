package model

type HomeText struct {
	Id int `json:"id"`
	Text string `json:"text"`
}

func GetHomeText() HomeText {
	var homeText HomeText
	db.Table("ml_index_text").First(&homeText)
	return homeText
}
