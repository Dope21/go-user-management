package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user-management/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string { "status": "ready!" }
		encoder := json.NewEncoder(w)
		encoder.Encode(response)
	})

	mux.HandleFunc("POST /users/register", handlers.RegisterUser)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	fmt.Println("Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}