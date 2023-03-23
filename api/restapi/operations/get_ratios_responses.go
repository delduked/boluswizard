// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"boluswizard/models"
)

// GetRatiosOKCode is the HTTP code returned for type GetRatiosOK
const GetRatiosOKCode int = 200

/*
GetRatiosOK Successful operation

swagger:response getRatiosOK
*/
type GetRatiosOK struct {

	/*
	  In: Body
	*/
	Payload *GetRatiosOKBody `json:"body,omitempty"`
}

// NewGetRatiosOK creates GetRatiosOK with default headers values
func NewGetRatiosOK() *GetRatiosOK {

	return &GetRatiosOK{}
}

// WithPayload adds the payload to the get ratios o k response
func (o *GetRatiosOK) WithPayload(payload *GetRatiosOKBody) *GetRatiosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ratios o k response
func (o *GetRatiosOK) SetPayload(payload *GetRatiosOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRatiosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetRatiosDefault Unsuccessful operation

swagger:response getRatiosDefault
*/
type GetRatiosDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetRatiosDefault creates GetRatiosDefault with default headers values
func NewGetRatiosDefault(code int) *GetRatiosDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRatiosDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ratios default response
func (o *GetRatiosDefault) WithStatusCode(code int) *GetRatiosDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ratios default response
func (o *GetRatiosDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ratios default response
func (o *GetRatiosDefault) WithPayload(payload *models.Response) *GetRatiosDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ratios default response
func (o *GetRatiosDefault) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRatiosDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
