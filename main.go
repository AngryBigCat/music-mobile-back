package main

import (
	"fmt"
	"music-mobile-back/router"
	"music-mobile-back/config"
)

func main() {
	r := router.Init()

	r.Run(fmt.Sprintf(":%s", config.LISTEN_PORT))
}
