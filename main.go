package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user-management/configs"
	"user-management/db"
	"user-management/handlers"
)

func main() {

	cfg := configs.LoadConfig()

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	db.RunMigrations(database)

	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string { "status": "ready!" }
		encoder := json.NewEncoder(w)
		encoder.Encode(response)
	})

	mux.HandleFunc("POST /users/register", handlers.RegisterUser)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", cfg.AppPort),
		Handler: mux,
	}

	fmt.Println("Server running on port " + cfg.AppPort)
	log.Fatal(server.ListenAndServe())
}