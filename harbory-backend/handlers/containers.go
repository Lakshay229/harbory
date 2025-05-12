package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"harbory-backend/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
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

func StartContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cli.Close()

	ctx := context.Background()
	err = cli.ContainerStart(ctx, containerID, container.StartOptions{})
	if err != nil {
		http.Error(w, "Failed to start container: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Container started successfully"))
}

func StopContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cli.Close()

	ctx := context.Background()
	timeoutSeconds := 30
	err = cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &timeoutSeconds})
	if err != nil {
		http.Error(w, "Failed to stop container: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Container stopped successfully"))
}

func DeleteContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, "Failed to create Docker client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cli.Close()

	ctx := context.Background()
	err = cli.ContainerRemove(ctx, containerID, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		http.Error(w, "Failed to delete container: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Container deleted successfully"))
}
