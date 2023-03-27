package tools

import (
	"boluswizard/restapi/operations"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func DeleteCookie(cookieName string, err error) middleware.ResponderFunc {
	authenticateResonse := func(rw http.ResponseWriter, pr runtime.Producer) {

		expiredCookie := &http.Cookie{
			Name:   cookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}

		response := operations.SignUpOKBody{
			Status: http.StatusUnauthorized,
			Error:  err,
		}
		http.SetCookie(rw, expiredCookie)

		rw.WriteHeader(http.StatusUnauthorized)

		pr.Produce(rw, response)
	}

	return authenticateResonse
}
