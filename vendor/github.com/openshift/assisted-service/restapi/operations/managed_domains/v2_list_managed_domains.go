// Code generated by go-swagger; DO NOT EDIT.

package managed_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2ListManagedDomainsHandlerFunc turns a function with the right signature into a v2 list managed domains handler
type V2ListManagedDomainsHandlerFunc func(V2ListManagedDomainsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2ListManagedDomainsHandlerFunc) Handle(params V2ListManagedDomainsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2ListManagedDomainsHandler interface for that can handle valid v2 list managed domains params
type V2ListManagedDomainsHandler interface {
	Handle(V2ListManagedDomainsParams, interface{}) middleware.Responder
}

// NewV2ListManagedDomains creates a new http.Handler for the v2 list managed domains operation
func NewV2ListManagedDomains(ctx *middleware.Context, handler V2ListManagedDomainsHandler) *V2ListManagedDomains {
	return &V2ListManagedDomains{Context: ctx, Handler: handler}
}

/* V2ListManagedDomains swagger:route GET /v2/domains managed_domains v2ListManagedDomains

List of managed DNS domains.

*/
type V2ListManagedDomains struct {
	Context *middleware.Context
	Handler V2ListManagedDomainsHandler
}

func (o *V2ListManagedDomains) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2ListManagedDomainsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}