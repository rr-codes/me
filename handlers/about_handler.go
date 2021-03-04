package handlers

import (
	"encoding/json"
	"me/models"
	"net/http"
)

func HandleAboutInfo(w http.ResponseWriter, _ *http.Request) {
	aboutInfo := models.GetAboutInfo()

	bytes, err := json.Marshal(aboutInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(bytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
