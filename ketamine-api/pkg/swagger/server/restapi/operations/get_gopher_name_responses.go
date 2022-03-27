// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetGopherNameOKCode is the HTTP code returned for type GetGopherNameOK
const GetGopherNameOKCode int = 200

/*GetGopherNameOK Returns the gopher.

swagger:response getGopherNameOK
*/
type GetGopherNameOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetGopherNameOK creates GetGopherNameOK with default headers values
func NewGetGopherNameOK() *GetGopherNameOK {

	return &GetGopherNameOK{}
}

// WithPayload adds the payload to the get gopher name o k response
func (o *GetGopherNameOK) WithPayload(payload io.ReadCloser) *GetGopherNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get gopher name o k response
func (o *GetGopherNameOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGopherNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
