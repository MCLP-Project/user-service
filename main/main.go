package main

import (
	"user-service/db"
	"user-service/models"
	"user-service/server"
)

func main() {
	server.RunServer(
		models.ServerConfig{
			Port: "9000",
		},
	)

	db.InitializeDB(models.DatabaseConfig{})
	

}
