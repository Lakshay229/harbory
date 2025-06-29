package main

import (
	"log"
	"net/http"

	"harbory-backend/internal/router"
)

func main() {
	r := router.SetupRouter()
	log.Println("Backend running on :8080")
	http.ListenAndServe(":8080", r)
}
