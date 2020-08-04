// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteProtectedEntityParams creates a new DeleteProtectedEntityParams object
// with the default values initialized.
func NewDeleteProtectedEntityParams() *DeleteProtectedEntityParams {
	var ()
	return &DeleteProtectedEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProtectedEntityParamsWithTimeout creates a new DeleteProtectedEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProtectedEntityParamsWithTimeout(timeout time.Duration) *DeleteProtectedEntityParams {
	var ()
	return &DeleteProtectedEntityParams{

		timeout: timeout,
	}
}

// NewDeleteProtectedEntityParamsWithContext creates a new DeleteProtectedEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProtectedEntityParamsWithContext(ctx context.Context) *DeleteProtectedEntityParams {
	var ()
	return &DeleteProtectedEntityParams{

		Context: ctx,
	}
}

// NewDeleteProtectedEntityParamsWithHTTPClient creates a new DeleteProtectedEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProtectedEntityParamsWithHTTPClient(client *http.Client) *DeleteProtectedEntityParams {
	var ()
	return &DeleteProtectedEntityParams{
		HTTPClient: client,
	}
}

/*DeleteProtectedEntityParams contains all the parameters to send to the API endpoint
for the delete protected entity operation typically these are written to a http.Request
*/
type DeleteProtectedEntityParams struct {

	/*ProtectedEntityID
	  The protected entity ID to retrieve info for

	*/
	ProtectedEntityID string
	/*Service
	  The service for the protected entity

	*/
	Service string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete protected entity params
func (o *DeleteProtectedEntityParams) WithTimeout(timeout time.Duration) *DeleteProtectedEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete protected entity params
func (o *DeleteProtectedEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete protected entity params
func (o *DeleteProtectedEntityParams) WithContext(ctx context.Context) *DeleteProtectedEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete protected entity params
func (o *DeleteProtectedEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete protected entity params
func (o *DeleteProtectedEntityParams) WithHTTPClient(client *http.Client) *DeleteProtectedEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete protected entity params
func (o *DeleteProtectedEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProtectedEntityID adds the protectedEntityID to the delete protected entity params
func (o *DeleteProtectedEntityParams) WithProtectedEntityID(protectedEntityID string) *DeleteProtectedEntityParams {
	o.SetProtectedEntityID(protectedEntityID)
	return o
}

// SetProtectedEntityID adds the protectedEntityId to the delete protected entity params
func (o *DeleteProtectedEntityParams) SetProtectedEntityID(protectedEntityID string) {
	o.ProtectedEntityID = protectedEntityID
}

// WithService adds the service to the delete protected entity params
func (o *DeleteProtectedEntityParams) WithService(service string) *DeleteProtectedEntityParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the delete protected entity params
func (o *DeleteProtectedEntityParams) SetService(service string) {
	o.Service = service
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProtectedEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param protectedEntityID
	if err := r.SetPathParam("protectedEntityID", o.ProtectedEntityID); err != nil {
		return err
	}

	// path param service
	if err := r.SetPathParam("service", o.Service); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}