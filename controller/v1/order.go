package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/AngryBigCat/music-mobile-back/model"
)

func GetOrder(c *gin.Context) {
	token := c.Query("token")
	orderNum := c.Query("ordernum")

	var order model.Order
	switch {
		case token != "":
			user, _ := model.ParseToken(token)
			order = model.GetOrder("phone", user.Phone)
		case orderNum != "":
			order = model.GetOrder("ordernum", orderNum)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"order": order,
	})
}

func GetOrders(c *gin.Context) {
	token := c.Query("token")
	user, _ := model.ParseToken(token)
	orders := model.GetOrders(user.Uid)

	c.JSON(200, gin.H{
		"code": 0,
		"orders": orders,
	})
}

var orderForm model.OrderForm
func PostOrder(c *gin.Context) {
	c.BindJSON(&orderForm)
	oForm, oNum := model.OrderStore(orderForm)
	c.JSON(200, gin.H{
		"code": 0,
		"total_price": oForm.Total_price,
		"ordernum": oNum,
	})
}