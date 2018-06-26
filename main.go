package main

import "music-mobile-back/router"

func main() {
	r := router.Init()

	r.Run()
}
