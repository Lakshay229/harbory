package api

import (
	"encoding/json"
	"io"
	"net/http"

	"harbory-backend/internal/domain"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ImageHandler struct {
	service  domain.ImageService
	upgrader websocket.Upgrader
}

func NewImageHandler(service domain.ImageService) *ImageHandler {
	return &ImageHandler{
		service: service,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *ImageHandler) GetImages(w http.ResponseWriter, r *http.Request) {
	images, err := h.service.List(r.Context(), true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (h *ImageHandler) InspectImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["id"]

	imageInfo, err := h.service.Inspect(r.Context(), imageID)
	if err != nil {
		http.Error(w, "Failed to inspect image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(imageInfo)
}

func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["id"]

	err := h.service.Delete(r.Context(), imageID)
	if err != nil {
		http.Error(w, "Failed to delete image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image deleted successfully"))
}

func (h *ImageHandler) PullImage(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ImageName string `json:"imageName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.Pull(r.Context(), req.ImageName, w)
	if err != nil {
		http.Error(w, "Failed to pull image: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ImageHandler) PullImageResp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["id"]

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	pr, pw := io.Pipe()

	go func() {
		err := h.service.Pull(r.Context(), imageID, pw)
		pw.Close()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Error pulling image: "+err.Error()))
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := pr.Read(buffer)
		if err != nil {
			if err != io.EOF {
				conn.WriteMessage(websocket.TextMessage, []byte("Error reading pull response: "+err.Error()))
			}
			break
		}

		conn.WriteMessage(websocket.TextMessage, buffer[:n])
	}
}
