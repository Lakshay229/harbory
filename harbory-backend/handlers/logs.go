package handlers

import (
	"net/http"

	"harbory-backend/utils"
	"log"
	"github.com/docker/docker/api/types/container"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},	
}

func GetContainerLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["id"]
	
	cli := utils.GetDockerClient()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	ctx := r.Context()
	reader, err := cli.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
		Tail:       "all",
	})
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Failed to get logs: "+err.Error()))
		return
	}
	defer reader.Close()

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			break
		}
		if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
			break
		}
	}
}
