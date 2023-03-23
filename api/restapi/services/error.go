package services

import (
	"boluswizard/models"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
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
func NewError(status int, err error) middleware.ResponderFunc {

	test := func(rw http.ResponseWriter, pr runtime.Producer) {
		rw.WriteHeader(status)
		response := &models.Response{
			Status: int64(status),
			Error:  err,
		}
		pr.Produce(rw, response)
	}

	return test

}
