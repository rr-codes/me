package api

import (
	"encoding/json"
	"me/models"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	aboutInfo := models.GetAboutInfo()
	bytes, _ := json.Marshal(aboutInfo)

	if _, err := w.Write(bytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
