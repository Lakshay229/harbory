package handlers

import (
	"encoding/json"
	"net/http"

	"harbory-backend/utils"

	"github.com/docker/docker/api/types/container"
)

func GetContainers(w http.ResponseWriter, r *http.Request) {
	cli := utils.GetDockerClient()
	containers, err := cli.ContainerList(r.Context(), container.ListOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(containers)
}
