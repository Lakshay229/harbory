package handlers

import (
	"encoding/json"
	"net/http"

	"harbory-backend/utils"
)

func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	cli := utils.GetDockerClient()
	info, err := cli.Info(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(info)
}