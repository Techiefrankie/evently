package main

import (
	"evently/config"
)

func main() {
	config.InitGorm()
	RegisterRoutes()
}
