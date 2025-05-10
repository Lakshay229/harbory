package router

import (
	"harbory-backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/containers", handlers.GetContainers).Methods("GET")
	r.HandleFunc("/api/images", handlers.GetImages).Methods("GET")
	r.HandleFunc("/api/system/info", handlers.GetSystemInfo).Methods("GET")

	return r
}
