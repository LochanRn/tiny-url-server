package handlers

import (
	"github.com/LochanRn/tiny-url-server/domain"
	v1 "github.com/LochanRn/tiny-url-server/gen/restapi/operations"
	"github.com/LochanRn/tiny-url-server/services/tinyurl"
	"github.com/pkg/errors"
)

type V1Handler struct {
	TinyUrlService *tinyurl.TinyURLService
}

func NewV1Handler(dbType domain.DBType) (*V1Handler, error) {
	tinyUrlService, err := tinyurl.NewTinyURLService(dbType)
	if err != nil {
		return nil, err
	}
	return &V1Handler{
		TinyUrlService: tinyUrlService,
	}, nil
}

func ConfigureV1Handlers(api *v1.TinyURLServerAPI) (*v1.TinyURLServerAPI, error) {
	h, err := NewV1Handler(domain.LocalStore)
	if err != nil {
		return nil, errors.Wrap(err, "error while creating new handler")
	}
	api.V1PingHandler = v1.V1PingHandlerFunc(h.Ping)
	api.V1TinyurlPostHandler = v1.V1TinyurlPostHandlerFunc(h.CreateTinyURL)
	api.V1DomainsShorternedHandler = v1.V1DomainsShorternedHandlerFunc(h.GetDomainsShortened)
	api.V1TinyurlRedirectHandler = v1.V1TinyurlRedirectHandlerFunc(h.Redirect)
	return api, nil
}
