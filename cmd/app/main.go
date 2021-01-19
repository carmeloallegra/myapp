package main

import (
	"fmt"
	"log"
	"myapp/cmd/app/config"
	"net/http"
	"time"

	"myapp/app/app/router"
	// "myapp/app/router"
)

func main() {
	appConf := config.AppConfig()

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      appRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

//Greet whatever
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
