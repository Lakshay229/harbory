package router

import (
	"harbory-backend/handlers"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	os.Setenv("DOCKER_API_VERSION", "1.47")
}

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/containers", handlers.GetContainers).Methods("GET")
	r.HandleFunc("/api/images", handlers.GetImages).Methods("GET")
	r.HandleFunc("/api/system/info", handlers.GetSystemInfo).Methods("GET")
	r.HandleFunc("/api/containers/{id}/logs", handlers.GetLogs)

	r.HandleFunc("/api/containers/{id}/start", handlers.StartContainer).Methods("POST")
	r.HandleFunc("/api/containers/{id}/stop", handlers.StopContainer).Methods("POST")
	r.HandleFunc("/api/containers/{id}/delete", handlers.DeleteContainer).Methods("DELETE")

	return r
}
