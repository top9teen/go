package main

import (
	"config"
	"route"
)

func main() {
	config.Init()

	// Route file
	route.Init()
}
