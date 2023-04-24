package main

import (
	"book/api"
	"book/config"
)

func main() {
	config.Init()
	api.StartServer()
}
