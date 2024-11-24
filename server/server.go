package server

import (
	"fmt"
	"log"
	"net/http"
	"user-service/constants"
	"user-service/models"
)

func RunServer(config models.Config) {
	http.HandleFunc(constants.HealthEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong"))
	})

	log.Printf("Starting server on :%v\n", config.ServerConfig.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.ServerConfig.Port), nil); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
	// if err := http.ListenAndServe(":9000", nil); err != nil {
	// 	log.Fatalf("could not start server: %s\n", err.Error())
	// }
}
