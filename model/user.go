package model

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"fmt"
	"github.com/goroom/aliyun_sms"
	"log"
	"music-mobile-back/config"
)

type User struct {
	Uid int `json:"id"`
	Phone string `json:"phone"`
	Password string
}

type Claims struct {
	Uid int
	Phone string
	Password string
	jwt.StandardClaims
}

var user User

func CheckUserLogin(username, password string) int {
	db.Table("ml_user").Where("phone=?", username).First(&user)
	if username != user.Phone || password != user.Password {
		return 0
	}
	return user.Uid
}


var jwtSecret = []byte("music")
func GenerateToken(uid int, username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
			uid,
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer : "music-mobile-back",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error)  {
	tokenClaims, e := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	claims, _ := tokenClaims.Claims.(*Claims)
	return claims, e
}

func SendRegisterCodeStoreRedis(phone string) bool {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	aliyunSms, err := aliyun_sms.
		NewAliyunSms(config.SIGNNAME, config.SENDCODE_TEMPLATE_CODE, config.ACCESS_KEY, config.ACCESS_SEC)
	if err != nil {
		log.Println(err)
		return false
	}
	err = aliyunSms.Send(phone, fmt.Sprintf(`{"code":"%s"}`, code))
	if err != nil {
		log.Println(err)
		return false
	}
	err = redisClient.Set("code:"+ phone, code, config.REDIS_SENDCODE_EX * time.Second).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func CheckCode(phone, code string) bool {
	c, _ := redisClient.Get("code:" + phone).Result()
	if c != code { return false }
	return true
}

func CheckReUsername(phone string) bool {
	db.Table("ml_user").Where("phone=?", phone).First(&user)
	if (User{}) != user {
		return false
	}
	return true
}

func CreateUser(phone, password string) int {
	user = User{Phone: phone, Password: password }
	db.Table("ml_user").Create(&user)
	db.Table("ml_user").Where("phone=?", phone).First(&user)
	return user.Uid
}
