package v1

import (
	"github.com/gin-gonic/gin"
	"music-mobile-back/model"
	"log"
)


func GetToken(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	userId := model.CheckUserLogin(username, password)
	if  userId == 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "用户名或密码错误",
		})
		return
	}

	token, err := model.GenerateToken(userId, username, password)
	if err!= nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "登陆成功",
		"token": token,
	})
}

func GetCode(c *gin.Context) {
	phone := c.Query("phone")
	if !model.SendRegisterCodeStoreRedis(phone) {
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "短信发送失败",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "短信发送成功",
		})
	}

}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := c.PostForm("code")

	if !model.CheckCode(username, code) {
		c.JSON(200, gin.H{
			"code": 1,
			"msg": "验证码错误",
		})
		return
	}
	if !model.CheckReUsername(username) {
		c.JSON(200, gin.H{
			"code": 2,
			"msg": "该手机号已被注册了",
		})
		return
	}
	userId := model.CreateUser(username, password)
	token, _ := model.GenerateToken(userId, username, password)
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "用户注册成功",
		"token": token,
	})
}
	