package services

import (
	"boluswizard/models"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func Healthcheck(rw http.ResponseWriter, pr runtime.Producer) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	response := &models.Health{
		Status:    200,
		Error:     nil,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	}
	pr.Produce(rw, response)
}
