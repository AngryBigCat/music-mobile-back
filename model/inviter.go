package model

type Inviter struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var inviters []Inviter

func GetInviters() []Inviter {
	db.Table("ml_inviter").Find(&inviters)
	return inviters
}