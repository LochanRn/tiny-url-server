// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V1TinyurlRedirectHandlerFunc turns a function with the right signature into a v1 tinyurl redirect handler
type V1TinyurlRedirectHandlerFunc func(V1TinyurlRedirectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn V1TinyurlRedirectHandlerFunc) Handle(params V1TinyurlRedirectParams) middleware.Responder {
	return fn(params)
}

// V1TinyurlRedirectHandler interface for that can handle valid v1 tinyurl redirect params
type V1TinyurlRedirectHandler interface {
	Handle(V1TinyurlRedirectParams) middleware.Responder
}

// NewV1TinyurlRedirect creates a new http.Handler for the v1 tinyurl redirect operation
func NewV1TinyurlRedirect(ctx *middleware.Context, handler V1TinyurlRedirectHandler) *V1TinyurlRedirect {
	return &V1TinyurlRedirect{Context: ctx, Handler: handler}
}

/*
	V1TinyurlRedirect swagger:route GET /v1/tinyurl/{tinyurl} tinyurl v1TinyurlRedirect

# Redirect to original URL

Redirect to original URL
*/
type V1TinyurlRedirect struct {
	Context *middleware.Context
	Handler V1TinyurlRedirectHandler
}

func (o *V1TinyurlRedirect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV1TinyurlRedirectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
