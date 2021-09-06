// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package plugins

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

// NewPluginsIPReservationsRelatedIPV6ReadParams creates a new PluginsIPReservationsRelatedIPV6ReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsIPReservationsRelatedIPV6ReadParams() *PluginsIPReservationsRelatedIPV6ReadParams {
	return &PluginsIPReservationsRelatedIPV6ReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsIPReservationsRelatedIPV6ReadParamsWithTimeout creates a new PluginsIPReservationsRelatedIPV6ReadParams object
// with the ability to set a timeout on a request.
func NewPluginsIPReservationsRelatedIPV6ReadParamsWithTimeout(timeout time.Duration) *PluginsIPReservationsRelatedIPV6ReadParams {
	return &PluginsIPReservationsRelatedIPV6ReadParams{
		timeout: timeout,
	}
}

// NewPluginsIPReservationsRelatedIPV6ReadParamsWithContext creates a new PluginsIPReservationsRelatedIPV6ReadParams object
// with the ability to set a context for a request.
func NewPluginsIPReservationsRelatedIPV6ReadParamsWithContext(ctx context.Context) *PluginsIPReservationsRelatedIPV6ReadParams {
	return &PluginsIPReservationsRelatedIPV6ReadParams{
		Context: ctx,
	}
}

// NewPluginsIPReservationsRelatedIPV6ReadParamsWithHTTPClient creates a new PluginsIPReservationsRelatedIPV6ReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsIPReservationsRelatedIPV6ReadParamsWithHTTPClient(client *http.Client) *PluginsIPReservationsRelatedIPV6ReadParams {
	return &PluginsIPReservationsRelatedIPV6ReadParams{
		HTTPClient: client,
	}
}

/* PluginsIPReservationsRelatedIPV6ReadParams contains all the parameters to send to the API endpoint
   for the plugins ip reservations related ipv6 read operation.

   Typically these are written to a http.Request.
*/
type PluginsIPReservationsRelatedIPV6ReadParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins ip reservations related ipv6 read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsIPReservationsRelatedIPV6ReadParams) WithDefaults() *PluginsIPReservationsRelatedIPV6ReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins ip reservations related ipv6 read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsIPReservationsRelatedIPV6ReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) WithTimeout(timeout time.Duration) *PluginsIPReservationsRelatedIPV6ReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) WithContext(ctx context.Context) *PluginsIPReservationsRelatedIPV6ReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) WithHTTPClient(client *http.Client) *PluginsIPReservationsRelatedIPV6ReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) WithID(id string) *PluginsIPReservationsRelatedIPV6ReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins ip reservations related ipv6 read params
func (o *PluginsIPReservationsRelatedIPV6ReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsIPReservationsRelatedIPV6ReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
