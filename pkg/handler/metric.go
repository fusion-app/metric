package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (handler *APIHandler) ListResourceMetrics(w http.ResponseWriter, r *http.Request) {
	metrics, err := handler.store.ListMetrics()
	if err != nil {
		log.Warningf("Failed to list metrics", err)
		responseJSON(Message{err.Error()}, w, http.StatusInternalServerError)
		return
	}
	responseJSON(metrics, w, http.StatusOK)
}
