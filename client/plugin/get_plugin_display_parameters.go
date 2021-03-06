// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPluginDisplayParams creates a new GetPluginDisplayParams object
// with the default values initialized.
func NewGetPluginDisplayParams() *GetPluginDisplayParams {
	var ()
	return &GetPluginDisplayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginDisplayParamsWithTimeout creates a new GetPluginDisplayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPluginDisplayParamsWithTimeout(timeout time.Duration) *GetPluginDisplayParams {
	var ()
	return &GetPluginDisplayParams{

		timeout: timeout,
	}
}

// NewGetPluginDisplayParamsWithContext creates a new GetPluginDisplayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPluginDisplayParamsWithContext(ctx context.Context) *GetPluginDisplayParams {
	var ()
	return &GetPluginDisplayParams{

		Context: ctx,
	}
}

// NewGetPluginDisplayParamsWithHTTPClient creates a new GetPluginDisplayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPluginDisplayParamsWithHTTPClient(client *http.Client) *GetPluginDisplayParams {
	var ()
	return &GetPluginDisplayParams{
		HTTPClient: client,
	}
}

/*GetPluginDisplayParams contains all the parameters to send to the API endpoint
for the get plugin display operation typically these are written to a http.Request
*/
type GetPluginDisplayParams struct {

	/*ID
	  the plugin id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get plugin display params
func (o *GetPluginDisplayParams) WithTimeout(timeout time.Duration) *GetPluginDisplayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin display params
func (o *GetPluginDisplayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin display params
func (o *GetPluginDisplayParams) WithContext(ctx context.Context) *GetPluginDisplayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin display params
func (o *GetPluginDisplayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin display params
func (o *GetPluginDisplayParams) WithHTTPClient(client *http.Client) *GetPluginDisplayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin display params
func (o *GetPluginDisplayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get plugin display params
func (o *GetPluginDisplayParams) WithID(id int64) *GetPluginDisplayParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get plugin display params
func (o *GetPluginDisplayParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginDisplayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
