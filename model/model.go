package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"music-mobile-back/config"
	"fmt"
)

var db *gorm.DB

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
}

func CloseDB() {
	defer db.Close()
}
