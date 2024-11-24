package main

import (
	"user-service/models"
	"user-service/server"
)

func main() {
	server.RunServer(models.Config{
		ServerConfig: models.ServerConfig{
			Port: "9000",
		},
	})
}
