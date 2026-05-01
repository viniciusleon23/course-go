package main

import (
	"course-go/config"
	"course-go/server"
)
func main() {
	config := config.LoadConfig()

	app := server.NewApp()
	app.RunServer(config.Port)

}