package v1

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/model"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/mozillazg/request"
)

type Form struct {
	Price float64
	Ordernum string
}

func Pay(c *gin.Context) {
	ordernum := c.PostForm("ordernum")
	state := c.PostForm("state")

	var url string
	switch state {
		case "wx":
			order := model.GetOrder("ordernum", ordernum)
			c := new(http.Client)
			req := request.NewRequest(c)
			req.Data = map[string]string{
				"price": fmt.Sprintf("%.2f", order.All_price),
				"ordernum": ordernum,
			}
			response, _ := req.Post("http://youbanquan.com/wxpay")
			defer response.Body.Close()

			result, e := ioutil.ReadAll(response.Body)
			if e != nil {
				panic(e)
			}
			url = fmt.Sprintf("http://youbanquan.com%s", result)
		case "al":
				order := model.GetOrder("ordernum", ordernum)
				url = fmt.Sprintf("http://youbanquan.com/zfb/%.2f/%s", order.All_price, order.Ordernum)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"state": state,
		"url": url,
	})
}
