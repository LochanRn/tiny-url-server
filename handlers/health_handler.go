package handlers

import (
	"github.com/LochanRn/tiny-url-api/gen/models"
	v1 "github.com/LochanRn/tiny-url-server/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (v *V1Handler) Ping(params v1.V1PingParams) middleware.Responder {
	return v1.NewV1PingOK().WithPayload(&models.V1Ping{
		Msg: "pong",
	})
}
