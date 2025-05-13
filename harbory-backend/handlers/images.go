package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"harbory-backend/utils"

	"github.com/docker/docker/api/types/image"
	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
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
	_, err = cli.ImageRemove(ctx, imageID, imagetypes.RemoveOptions{
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

func InspectImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cli.Close()

	ctx := context.Background()
	inspect, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		http.Error(w, "Failed to inspect image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(inspect)
}
