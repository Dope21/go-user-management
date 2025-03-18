package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/configs"
	"user-management/db"
	"user-management/routes"
)

func main() {

	cfg := configs.LoadConfig()

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	db.RunMigrations(database)

	defer database.Close()

	mux := routes.Router()

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.AppPort),
		Handler: mux,
	}

	fmt.Println("Server running on port " + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}