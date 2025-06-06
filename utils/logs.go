package utils

import (
	"log"
	"net/http"
)

func LogError(r *http.Request, message string) {
	log.Printf("| ERROR | %s | %s", r.URL.String(), message)
}

func LogInfo(r *http.Request, message string) {
	log.Printf("| INFO | %s | %s", r.URL.String(), message)
}