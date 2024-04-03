package handlers

import (
	v1 "github.com/LochanRn/tiny-url-server/gen/restapi/operations"
)

type V1Handler struct {
}

func NewV1Handler() *V1Handler {
	return &V1Handler{}
}

func ConfigureV1Handlers(api *v1.TinyURLServerAPI) (*v1.TinyURLServerAPI, error) {
	h := NewV1Handler()
	api.V1PingHandler = v1.V1PingHandlerFunc(h.Ping)
	return api, nil
}
