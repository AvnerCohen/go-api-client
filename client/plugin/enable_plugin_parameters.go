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

// NewEnablePluginParams creates a new EnablePluginParams object
// with the default values initialized.
func NewEnablePluginParams() *EnablePluginParams {
	var ()
	return &EnablePluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnablePluginParamsWithTimeout creates a new EnablePluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnablePluginParamsWithTimeout(timeout time.Duration) *EnablePluginParams {
	var ()
	return &EnablePluginParams{

		timeout: timeout,
	}
}

// NewEnablePluginParamsWithContext creates a new EnablePluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnablePluginParamsWithContext(ctx context.Context) *EnablePluginParams {
	var ()
	return &EnablePluginParams{

		Context: ctx,
	}
}

// NewEnablePluginParamsWithHTTPClient creates a new EnablePluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnablePluginParamsWithHTTPClient(client *http.Client) *EnablePluginParams {
	var ()
	return &EnablePluginParams{
		HTTPClient: client,
	}
}

/*EnablePluginParams contains all the parameters to send to the API endpoint
for the enable plugin operation typically these are written to a http.Request
*/
type EnablePluginParams struct {

	/*ID
	  the plugin id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable plugin params
func (o *EnablePluginParams) WithTimeout(timeout time.Duration) *EnablePluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable plugin params
func (o *EnablePluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable plugin params
func (o *EnablePluginParams) WithContext(ctx context.Context) *EnablePluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable plugin params
func (o *EnablePluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable plugin params
func (o *EnablePluginParams) WithHTTPClient(client *http.Client) *EnablePluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable plugin params
func (o *EnablePluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the enable plugin params
func (o *EnablePluginParams) WithID(id int64) *EnablePluginParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enable plugin params
func (o *EnablePluginParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnablePluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
