package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io"
	"context" 

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"

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

func PullImageResp(w http.ResponseWriter, r *http.Request) {
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

func PullImage(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Image string `json:"image"`
		Tag   string `json:"tag"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if request.Image == "" {
		http.Error(w, "Image name is required", http.StatusBadRequest)
		return
	}

	if request.Tag == "" {
		request.Tag = "latest"
	}

	imageName := request.Image + ":" + request.Tag

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cli.Close()
	ctx := context.Background()
	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		http.Error(w, "Failed to pull image: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	io.Copy(io.Discard, out)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image pulled successfully"))
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cli.Close()

	ctx := context.Background()
	_, err = cli.ImageRemove(ctx, imageID, image.RemoveOptions{
		Force:         true,
		PruneChildren: true,
	})

	if err != nil {
		http.Error(w, "Failed to delete image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image deleted successfully"))
}