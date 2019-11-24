package handler

import (
	"encoding/json"
	"github.com/fusion-app/metric/pkg/store"
	"net/http"
)

type APIHandler struct {
	frontDir string
	store    store.Store
}

func NewAPIHandler(frontDir string) (*APIHandler, error) {
	localStore := &store.LocalStore{}

	apiHandler := &APIHandler{
		frontDir: frontDir,
		store:    localStore,
	}
	return apiHandler, nil
}

type Message struct {
	Message string `json:"message"`
}

func responseJSON(body interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
