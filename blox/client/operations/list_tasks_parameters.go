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

// NewListTasksParams creates a new ListTasksParams object
// with the default values initialized.
func NewListTasksParams() *ListTasksParams {
	var ()
	return &ListTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksParamsWithTimeout creates a new ListTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTasksParamsWithTimeout(timeout time.Duration) *ListTasksParams {
	var ()
	return &ListTasksParams{

		timeout: timeout,
	}
}

// NewListTasksParamsWithContext creates a new ListTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTasksParamsWithContext(ctx context.Context) *ListTasksParams {
	var ()
	return &ListTasksParams{

		Context: ctx,
	}
}

// NewListTasksParamsWithHTTPClient creates a new ListTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTasksParamsWithHTTPClient(client *http.Client) *ListTasksParams {
	var ()
	return &ListTasksParams{
		HTTPClient: client,
	}
}

/*ListTasksParams contains all the parameters to send to the API endpoint
for the list tasks operation typically these are written to a http.Request
*/
type ListTasksParams struct {

	/*Cluster
	  Cluster name or ARN to filter tasks by

	*/
	Cluster *string
	/*StartedBy
	  StartedBy to filter tasks by

	*/
	StartedBy *string
	/*Status
	  Status to filter tasks by

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) WithTimeout(timeout time.Duration) *ListTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tasks params
func (o *ListTasksParams) WithContext(ctx context.Context) *ListTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tasks params
func (o *ListTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tasks params
func (o *ListTasksParams) WithHTTPClient(client *http.Client) *ListTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tasks params
func (o *ListTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCluster adds the cluster to the list tasks params
func (o *ListTasksParams) WithCluster(cluster *string) *ListTasksParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the list tasks params
func (o *ListTasksParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithStartedBy adds the startedBy to the list tasks params
func (o *ListTasksParams) WithStartedBy(startedBy *string) *ListTasksParams {
	o.SetStartedBy(startedBy)
	return o
}

// SetStartedBy adds the startedBy to the list tasks params
func (o *ListTasksParams) SetStartedBy(startedBy *string) {
	o.StartedBy = startedBy
}

// WithStatus adds the status to the list tasks params
func (o *ListTasksParams) WithStatus(status *string) *ListTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list tasks params
func (o *ListTasksParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string
		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {
			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}

	}

	if o.StartedBy != nil {

		// query param startedBy
		var qrStartedBy string
		if o.StartedBy != nil {
			qrStartedBy = *o.StartedBy
		}
		qStartedBy := qrStartedBy
		if qStartedBy != "" {
			if err := r.SetQueryParam("startedBy", qStartedBy); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
