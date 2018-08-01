package model

import (
	"time"
	"math/rand"
	"fmt"
)

type Order struct {
	Oid int `json:"oid"`
	Name string `json:"name"`
	Site string `json:"site"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Gsname string `json:"gsname"`
	Gsnum string `json:"gsnum"`
	Gsiphone string `json:"gsiphone"`
	Text string `json:"text"`
	Goodsname string `json:"goodsname"`
	Link string `json:"link"`
	Pname string `json:"pname"`
	Otime string `json:"otime"`
	Ordernum string `json:"ordernum"`
	O_price float64 `json:"o_price"`
	All_price float64
	Year int `json:"year"`
	City int `json:"city"`
	O_num string `json:"o_num"`
	Id int
	Userid int
	Inviter_kefu string
	Music_name string `json:"music_name"`
	Pay int `json:"pay"`
}


type OrderForm struct {
	Info_company string
	Info_name string
	Info_email string
	Info_phone string
	Info_remark string
	Info_inviter string
	Cert_company string
	Cert_num string
	Cert_address string
	Cert_phone string
	Cert_sign string
	Cert_link string
	Musics []OrderFormMusic
	Total_price float64
	Token string
}

type OrderFormMusic struct {
	Music_id int
	Name string
	Commerces_price float64
	Nocommerces_price float64
	Real_price float64
	Commerces_text int
	Region_text int
	Limit_text int
	Checked bool
	Number string
}


var order Order
func GetOrder(field, value string) Order {
	db.Table("ml_order").Where(field + "=?", value).Order("otime desc").First(&order)
	return order
}

func GetOrders(id int) []Order {
	var orders []Order
	db.Table("ml_order").Where("userid=?", id).Order("otime desc").Find(&orders)
	return orders
}

func OrderStore(orderForm OrderForm) (OrderForm, string) {
	user, _ := ParseToken(orderForm.Token)

	orderNum := generateOrderNum()

	for _, music := range orderForm.Musics {
		order := Order{
			Name: orderForm.Info_company,
			Site: orderForm.Cert_address,
			Email: orderForm.Info_email,
			Phone: orderForm.Info_phone,
			Gsname: orderForm.Cert_company,
			Gsnum: orderForm.Cert_num,
			Gsiphone: orderForm.Cert_phone,
			Text: orderForm.Info_remark,
			Goodsname: orderForm.Cert_sign,
			Link: orderForm.Cert_link,
			Pname: orderForm.Info_name,
			Otime: time.Now().Format("2006-01-02 15:04:05"),
			Ordernum: orderNum,
			O_price: music.Real_price,
			All_price: orderForm.Total_price,
			Year: music.Limit_text,
			City: music.Region_text,
			O_num: music.Number,
			Id: music.Music_id,
			Userid: user.Uid,
			Inviter_kefu: orderForm.Info_inviter,
			Music_name: music.Name,
		}
		db.Table("ml_order").Create(&order)
	}
	return orderForm, orderNum
}


func generateOrderNum() string {
	var orderOther Order
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderNum := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	db.Table("ml_order").Where("ordernum=?", orderNum).First(&orderOther)
	if orderOther != (Order{}) {
		return generateOrderNum()
	}
	return orderNum
}