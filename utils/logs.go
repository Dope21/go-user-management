package utils

import (
	"log"
	"net/http"
)

func LogError(r *http.Request, err error) {
	log.Printf("| ERROR | %s | %s", r.URL.String(), err.Error())
}

func LogInfo(r *http.Request, message string) {
	log.Printf("| INFO | %s | %s", r.URL.String(), message)
}