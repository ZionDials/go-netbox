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

package dcim

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
	"github.com/go-openapi/swag"
)

// NewDcimFrontPortTemplatesDeleteParams creates a new DcimFrontPortTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesDeleteParams() *DcimFrontPortTemplatesDeleteParams {
	return &DcimFrontPortTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesDeleteParamsWithTimeout creates a new DcimFrontPortTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesDeleteParams {
	return &DcimFrontPortTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesDeleteParamsWithContext creates a new DcimFrontPortTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesDeleteParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesDeleteParams {
	return &DcimFrontPortTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesDeleteParamsWithHTTPClient creates a new DcimFrontPortTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesDeleteParams {
	return &DcimFrontPortTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim front port templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this front port template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesDeleteParams) WithDefaults() *DcimFrontPortTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) WithID(id int64) *DcimFrontPortTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates delete params
func (o *DcimFrontPortTemplatesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
