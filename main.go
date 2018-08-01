package main

import (
	"github.com/AngryBigCat/music-mobile-back/router"
	"fmt"
	"github.com/AngryBigCat/music-mobile-back/config"
)

func main() {
	r := router.Init()
	r.Run(fmt.Sprintf(":%s", config.LISTEN_PORT))
}
