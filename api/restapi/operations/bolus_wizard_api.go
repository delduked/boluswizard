// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewBolusWizardAPI creates a new BolusWizard instance
func NewBolusWizardAPI(spec *loads.Document) *BolusWizardAPI {
	return &BolusWizardAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CorrectionRangeHandler: CorrectionRangeHandlerFunc(func(params CorrectionRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation CorrectionRange has not yet been implemented")
		}),
		CreateCorrectionsHandler: CreateCorrectionsHandlerFunc(func(params CreateCorrectionsParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateCorrections has not yet been implemented")
		}),
		CreateDurationHandler: CreateDurationHandlerFunc(func(params CreateDurationParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateDuration has not yet been implemented")
		}),
		CreateISFsHandler: CreateISFsHandlerFunc(func(params CreateISFsParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateISFs has not yet been implemented")
		}),
		CreateRatiosHandler: CreateRatiosHandlerFunc(func(params CreateRatiosParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateRatios has not yet been implemented")
		}),
		CreateTargetsHandler: CreateTargetsHandlerFunc(func(params CreateTargetsParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateTargets has not yet been implemented")
		}),
		GetCorrectionsHandler: GetCorrectionsHandlerFunc(func(params GetCorrectionsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetCorrections has not yet been implemented")
		}),
		GetDurationHandler: GetDurationHandlerFunc(func(params GetDurationParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDuration has not yet been implemented")
		}),
		GetISFsHandler: GetISFsHandlerFunc(func(params GetISFsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetISFs has not yet been implemented")
		}),
		GetRatiosHandler: GetRatiosHandlerFunc(func(params GetRatiosParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRatios has not yet been implemented")
		}),
		GetTargetsHandler: GetTargetsHandlerFunc(func(params GetTargetsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTargets has not yet been implemented")
		}),
		HealthCheckHandler: HealthCheckHandlerFunc(func(params HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthCheck has not yet been implemented")
		}),
		NewCorrectionHandler: NewCorrectionHandlerFunc(func(params NewCorrectionParams) middleware.Responder {
			return middleware.NotImplemented("operation NewCorrection has not yet been implemented")
		}),
		SignInHandler: SignInHandlerFunc(func(params SignInParams) middleware.Responder {
			return middleware.NotImplemented("operation SignIn has not yet been implemented")
		}),
		SignUpHandler: SignUpHandlerFunc(func(params SignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation SignUp has not yet been implemented")
		}),
	}
}

/*BolusWizardAPI the bolus wizard API */
type BolusWizardAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CorrectionRangeHandler sets the operation handler for the correction range operation
	CorrectionRangeHandler CorrectionRangeHandler
	// CreateCorrectionsHandler sets the operation handler for the create corrections operation
	CreateCorrectionsHandler CreateCorrectionsHandler
	// CreateDurationHandler sets the operation handler for the create duration operation
	CreateDurationHandler CreateDurationHandler
	// CreateISFsHandler sets the operation handler for the create i s fs operation
	CreateISFsHandler CreateISFsHandler
	// CreateRatiosHandler sets the operation handler for the create ratios operation
	CreateRatiosHandler CreateRatiosHandler
	// CreateTargetsHandler sets the operation handler for the create targets operation
	CreateTargetsHandler CreateTargetsHandler
	// GetCorrectionsHandler sets the operation handler for the get corrections operation
	GetCorrectionsHandler GetCorrectionsHandler
	// GetDurationHandler sets the operation handler for the get duration operation
	GetDurationHandler GetDurationHandler
	// GetISFsHandler sets the operation handler for the get i s fs operation
	GetISFsHandler GetISFsHandler
	// GetRatiosHandler sets the operation handler for the get ratios operation
	GetRatiosHandler GetRatiosHandler
	// GetTargetsHandler sets the operation handler for the get targets operation
	GetTargetsHandler GetTargetsHandler
	// HealthCheckHandler sets the operation handler for the health check operation
	HealthCheckHandler HealthCheckHandler
	// NewCorrectionHandler sets the operation handler for the new correction operation
	NewCorrectionHandler NewCorrectionHandler
	// SignInHandler sets the operation handler for the sign in operation
	SignInHandler SignInHandler
	// SignUpHandler sets the operation handler for the sign up operation
	SignUpHandler SignUpHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BolusWizardAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BolusWizardAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BolusWizardAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BolusWizardAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BolusWizardAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BolusWizardAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BolusWizardAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BolusWizardAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BolusWizardAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BolusWizardAPI
func (o *BolusWizardAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CorrectionRangeHandler == nil {
		unregistered = append(unregistered, "CorrectionRangeHandler")
	}
	if o.CreateCorrectionsHandler == nil {
		unregistered = append(unregistered, "CreateCorrectionsHandler")
	}
	if o.CreateDurationHandler == nil {
		unregistered = append(unregistered, "CreateDurationHandler")
	}
	if o.CreateISFsHandler == nil {
		unregistered = append(unregistered, "CreateISFsHandler")
	}
	if o.CreateRatiosHandler == nil {
		unregistered = append(unregistered, "CreateRatiosHandler")
	}
	if o.CreateTargetsHandler == nil {
		unregistered = append(unregistered, "CreateTargetsHandler")
	}
	if o.GetCorrectionsHandler == nil {
		unregistered = append(unregistered, "GetCorrectionsHandler")
	}
	if o.GetDurationHandler == nil {
		unregistered = append(unregistered, "GetDurationHandler")
	}
	if o.GetISFsHandler == nil {
		unregistered = append(unregistered, "GetISFsHandler")
	}
	if o.GetRatiosHandler == nil {
		unregistered = append(unregistered, "GetRatiosHandler")
	}
	if o.GetTargetsHandler == nil {
		unregistered = append(unregistered, "GetTargetsHandler")
	}
	if o.HealthCheckHandler == nil {
		unregistered = append(unregistered, "HealthCheckHandler")
	}
	if o.NewCorrectionHandler == nil {
		unregistered = append(unregistered, "NewCorrectionHandler")
	}
	if o.SignInHandler == nil {
		unregistered = append(unregistered, "SignInHandler")
	}
	if o.SignUpHandler == nil {
		unregistered = append(unregistered, "SignUpHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BolusWizardAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BolusWizardAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *BolusWizardAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BolusWizardAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BolusWizardAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BolusWizardAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the bolus wizard API
func (o *BolusWizardAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BolusWizardAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/CorrectionRange"] = NewCorrectionRange(o.context, o.CorrectionRangeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/wizard/corrections"] = NewCreateCorrections(o.context, o.CreateCorrectionsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/wizard/Duration"] = NewCreateDuration(o.context, o.CreateDurationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/wizard/ISF"] = NewCreateISFs(o.context, o.CreateISFsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/wizard/ratios"] = NewCreateRatios(o.context, o.CreateRatiosHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/wizard/Targets"] = NewCreateTargets(o.context, o.CreateTargetsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/corrections"] = NewGetCorrections(o.context, o.GetCorrectionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/Duration"] = NewGetDuration(o.context, o.GetDurationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/ISF"] = NewGetISFs(o.context, o.GetISFsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/ratios"] = NewGetRatios(o.context, o.GetRatiosHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/Targets"] = NewGetTargets(o.context, o.GetTargetsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = NewHealthCheck(o.context, o.HealthCheckHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/wizard/NewCorrection"] = NewNewCorrection(o.context, o.NewCorrectionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/SignIn"] = NewSignIn(o.context, o.SignInHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/SignUp"] = NewSignUp(o.context, o.SignUpHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BolusWizardAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BolusWizardAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BolusWizardAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BolusWizardAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BolusWizardAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
