package handlers

import (
	"encoding/json"
	"net/http"

	"harbory-backend/utils"

	"github.com/docker/docker/api/types/image"
)

func GetImages(w http.ResponseWriter, r *http.Request) {
	cli := utils.GetDockerClient()
	images, err := cli.ImageList(r.Context(), image.ListOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(images)
}
