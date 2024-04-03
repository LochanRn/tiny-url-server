// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V1PingHandlerFunc turns a function with the right signature into a v1 ping handler
type V1PingHandlerFunc func(V1PingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn V1PingHandlerFunc) Handle(params V1PingParams) middleware.Responder {
	return fn(params)
}

// V1PingHandler interface for that can handle valid v1 ping params
type V1PingHandler interface {
	Handle(V1PingParams) middleware.Responder
}

// NewV1Ping creates a new http.Handler for the v1 ping operation
func NewV1Ping(ctx *middleware.Context, handler V1PingHandler) *V1Ping {
	return &V1Ping{Context: ctx, Handler: handler}
}

/*
	V1Ping swagger:route GET /v1/ping troubleshoot v1Ping

# Ping Service

Ping Service
*/
type V1Ping struct {
	Context *middleware.Context
	Handler V1PingHandler
}

func (o *V1Ping) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV1PingParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
