package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type LogsHandler struct {
	containerHandler *ContainerHandler
	upgrader         websocket.Upgrader
}

func NewLogsHandler(containerHandler *ContainerHandler) *LogsHandler {
	return &LogsHandler{
		containerHandler: containerHandler,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *LogsHandler) GetContainerLogs(w http.ResponseWriter, r *http.Request) {
	h.containerHandler.GetContainerLogs(w, r)
}
