package handlers

import (
	v1 "github.com/LochanRn/tiny-url-server/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (v *V1Handler) Tinyurl(params v1.V1TinyurlPostParams) middleware.Responder {

	// v1.NewV1TinyurlPostOK().WithPayload()
	// Write your code here
	return middleware.NotImplemented("operation v1.Tinyurl has not yet been implemented")
}
