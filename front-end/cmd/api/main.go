package main

import (
	"fmt"
	_ "front-end/docs"
	"log"
	"net/http"
)

const (
	webPort           = "8080"
	shortenerHost     = "test-microservice"
	shortenerGRPCPort = "50001"
)

// @title test task FrontEnd
// @version 1.0

// @host localhost:8080
// @BasePath /
func main() {

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: getRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
