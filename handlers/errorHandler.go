package handlers

import (
	"encoding/json"
	"jarvis-backend/models"
	"net/http"
)

func ErrorResponse(rw *http.ResponseWriter, codigoError int, mensaje string) {
	(*rw).WriteHeader(http.StatusNotFound)
	(*rw).Header().Set("Content-Type", "Application/json")
	jsonResponse, _ := json.Marshal((models.Error(codigoError, mensaje)))
	(*rw).Write(jsonResponse)
}
