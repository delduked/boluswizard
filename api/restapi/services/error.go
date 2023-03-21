package services

import (
	"boluswizard/models"
	"net/http"

	"github.com/go-openapi/runtime"
)

func Error(rw http.ResponseWriter, pr runtime.Producer) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)
	response := &models.Response{
		Status: http.StatusInternalServerError,
		Error:  "Internal Server Error",
	}
	pr.Produce(rw, response)
}
