package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/configs"
	"user-management/repository"
	"user-management/routes"
)

func main() {

	cfg := configs.LoadConfig()

	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	repository.RunMigrations()

	defer db.Close()

	mux := routes.Router()

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.AppPort),
		Handler: mux,
	}

	fmt.Println("Server running on port " + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}