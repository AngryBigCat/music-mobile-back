package main

import (
	"music-mobile-back/router"
	"fmt"
	"music-mobile-back/config"
)

func main() {
	r := router.Init()
	r.Run(fmt.Sprintf(":%s", config.LISTEN_PORT))
}
