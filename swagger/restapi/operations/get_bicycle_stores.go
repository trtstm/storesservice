// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBicycleStoresHandlerFunc turns a function with the right signature into a get bicycle stores handler
type GetBicycleStoresHandlerFunc func(GetBicycleStoresParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBicycleStoresHandlerFunc) Handle(params GetBicycleStoresParams) middleware.Responder {
	return fn(params)
}

// GetBicycleStoresHandler interface for that can handle valid get bicycle stores params
type GetBicycleStoresHandler interface {
	Handle(GetBicycleStoresParams) middleware.Responder
}

// NewGetBicycleStores creates a new http.Handler for the get bicycle stores operation
func NewGetBicycleStores(ctx *middleware.Context, handler GetBicycleStoresHandler) *GetBicycleStores {
	return &GetBicycleStores{Context: ctx, Handler: handler}
}

/*GetBicycleStores swagger:route GET /bicyclestores getBicycleStores

GetBicycleStores get bicycle stores API

*/
type GetBicycleStores struct {
	Context *middleware.Context
	Handler GetBicycleStoresHandler
}

func (o *GetBicycleStores) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBicycleStoresParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}