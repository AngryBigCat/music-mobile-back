package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"music-mobile-back/config"
	"fmt"
	"github.com/go-redis/redis"
)

var db *gorm.DB

var redisClient *redis.Client

// 初始化数据库
func init() {
	var err error
	db, err = gorm.Open(config.DBTYPE, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DBUSER,
			config.DBPASS,
			config.DBHOST,
			config.DBPORT,
			config.DBNAME,
		))
	if err != nil {
		log.Println(err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST,
		Password: config.REDIS_PASS, // no password set
		DB:       config.REDIS_DB,  // use default DB
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Println(pong, err)
	}

}
