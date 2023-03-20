// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"boluswizard/restapi/operations"
)

//go:generate swagger generate server --target ../../api --name BolusWizard --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.BolusWizardAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BolusWizardAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CorrectionRangeHandler == nil {
		api.CorrectionRangeHandler = operations.CorrectionRangeHandlerFunc(func(params operations.CorrectionRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CorrectionRange has not yet been implemented")
		})
	}
	if api.CreateCorrectionsHandler == nil {
		api.CreateCorrectionsHandler = operations.CreateCorrectionsHandlerFunc(func(params operations.CreateCorrectionsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateCorrections has not yet been implemented")
		})
	}
	if api.CreateDurationHandler == nil {
		api.CreateDurationHandler = operations.CreateDurationHandlerFunc(func(params operations.CreateDurationParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateDuration has not yet been implemented")
		})
	}
	if api.CreateISFsHandler == nil {
		api.CreateISFsHandler = operations.CreateISFsHandlerFunc(func(params operations.CreateISFsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateISFs has not yet been implemented")
		})
	}
	if api.CreateRatiosHandler == nil {
		api.CreateRatiosHandler = operations.CreateRatiosHandlerFunc(func(params operations.CreateRatiosParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateRatios has not yet been implemented")
		})
	}
	if api.CreateTargetsHandler == nil {
		api.CreateTargetsHandler = operations.CreateTargetsHandlerFunc(func(params operations.CreateTargetsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateTargets has not yet been implemented")
		})
	}
	if api.GetCorrectionsHandler == nil {
		api.GetCorrectionsHandler = operations.GetCorrectionsHandlerFunc(func(params operations.GetCorrectionsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCorrections has not yet been implemented")
		})
	}
	if api.GetDurationHandler == nil {
		api.GetDurationHandler = operations.GetDurationHandlerFunc(func(params operations.GetDurationParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetDuration has not yet been implemented")
		})
	}
	if api.GetISFsHandler == nil {
		api.GetISFsHandler = operations.GetISFsHandlerFunc(func(params operations.GetISFsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetISFs has not yet been implemented")
		})
	}
	if api.HealthCheckHandler == nil {
		api.HealthCheckHandler = operations.HealthCheckHandlerFunc(func(params operations.HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.HealthCheck has not yet been implemented")
		})
	}
	if api.NewCorrectionHandler == nil {
		api.NewCorrectionHandler = operations.NewCorrectionHandlerFunc(func(params operations.NewCorrectionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.NewCorrection has not yet been implemented")
		})
	}
	if api.SignInHandler == nil {
		api.SignInHandler = operations.SignInHandlerFunc(func(params operations.SignInParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.SignIn has not yet been implemented")
		})
	}
	if api.SignUpHandler == nil {
		api.SignUpHandler = operations.SignUpHandlerFunc(func(params operations.SignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.SignUp has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
