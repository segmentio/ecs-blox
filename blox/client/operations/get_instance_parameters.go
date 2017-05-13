package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstanceParams creates a new GetInstanceParams object
// with the default values initialized.
func NewGetInstanceParams() *GetInstanceParams {
	var ()
	return &GetInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceParamsWithTimeout creates a new GetInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstanceParamsWithTimeout(timeout time.Duration) *GetInstanceParams {
	var ()
	return &GetInstanceParams{

		timeout: timeout,
	}
}

// NewGetInstanceParamsWithContext creates a new GetInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstanceParamsWithContext(ctx context.Context) *GetInstanceParams {
	var ()
	return &GetInstanceParams{

		Context: ctx,
	}
}

// NewGetInstanceParamsWithHTTPClient creates a new GetInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstanceParamsWithHTTPClient(client *http.Client) *GetInstanceParams {
	var ()
	return &GetInstanceParams{
		HTTPClient: client,
	}
}

/*GetInstanceParams contains all the parameters to send to the API endpoint
for the get instance operation typically these are written to a http.Request
*/
type GetInstanceParams struct {

	/*Arn
	  ARN of the instance to fetch

	*/
	Arn string
	/*Cluster
	  Cluster name of the instance to fetch

	*/
	Cluster string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instance params
func (o *GetInstanceParams) WithTimeout(timeout time.Duration) *GetInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance params
func (o *GetInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance params
func (o *GetInstanceParams) WithContext(ctx context.Context) *GetInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance params
func (o *GetInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance params
func (o *GetInstanceParams) WithHTTPClient(client *http.Client) *GetInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance params
func (o *GetInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArn adds the arn to the get instance params
func (o *GetInstanceParams) WithArn(arn string) *GetInstanceParams {
	o.SetArn(arn)
	return o
}

// SetArn adds the arn to the get instance params
func (o *GetInstanceParams) SetArn(arn string) {
	o.Arn = arn
}

// WithCluster adds the cluster to the get instance params
func (o *GetInstanceParams) WithCluster(cluster string) *GetInstanceParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the get instance params
func (o *GetInstanceParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param arn
	if err := r.SetPathParam("arn", o.Arn); err != nil {
		return err
	}

	// path param cluster
	if err := r.SetPathParam("cluster", o.Cluster); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
