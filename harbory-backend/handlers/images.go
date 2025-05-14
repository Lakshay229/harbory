package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/docker/docker/api/types/image"
	"github.com/gorilla/websocket"
	"harbory-backend/utils"
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

// Get layers of an image
func InspectImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["id"]
	cli := utils.GetDockerClient()

	inspects, err := cli.ImageInspect(r.Context(), imageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(inspects)
}

func RemoveImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["id"]
	cli := utils.GetDockerClient()

	options := image.RemoveOptions{
		Force:         true, // Force removal of the image, even if it is being used by stopped containers or has dependent child images.
		PruneChildren: true, //Do not delete untagged parent images.
	}

	remove, err := cli.ImageRemove(r.Context(), imageId, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(remove)
}

func PullImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageName := vars["id"]
	cli := utils.GetDockerClient()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	options := image.PullOptions{
		All:      false, // Pull all images
		Platform: "",    // Specify the platform to pull
	}

	pullResp, err := cli.ImagePull(r.Context(), imageName, options)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Failed to get logs: "+err.Error()))
		return
	}
	defer pullResp.Close()

	dec := json.NewDecoder(pullResp)
	for {
		var msg map[string]interface{}
		if err := dec.Decode(&msg); err != nil {
			break
		}

		data, _ := json.Marshal(msg)
		conn.WriteMessage(websocket.TextMessage, data)
	}

}
