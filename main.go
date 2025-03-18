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

	r := routes.Router()

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.AppPort),
		Handler: r,
	}

	fmt.Println("Server running on port " + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}