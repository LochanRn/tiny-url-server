package handlers

import (
	"github.com/LochanRn/tiny-url-api/gen/models"
	v1 "github.com/LochanRn/tiny-url-server/gen/restapi/operations"
	"github.com/LochanRn/tiny-url-server/handlers/converters"
	"github.com/go-openapi/runtime/middleware"
)

func (v *V1Handler) CreateTinyURL(params v1.V1TinyurlPostParams) middleware.Responder {
	tinyURL, err := v.TinyUrlService.CreateTinyURL(params.Body.URL)
	if err != nil {
		return v1.NewV1DomainsShorternedInternalServerError().WithPayload(&models.V1InternalServerError{
			Error: err.Error(),
		})
	}
	return v1.NewV1TinyurlPostOK().WithPayload(&v1.V1TinyurlPostOKBody{
		Tinyurl: tinyURL.TinyURL,
		URL:     tinyURL.URL,
		// CreationTimestamp: tinyURL.CreationTimeStamp,
	})
}

func (v *V1Handler) GetDomainsShortened(params v1.V1DomainsShorternedParams) middleware.Responder {
	domainCounter, err := v.TinyUrlService.GetDomainsShortened()
	if err != nil {
		return v1.NewV1DomainsShorternedInternalServerError().WithPayload(&models.V1InternalServerError{
			Error: err.Error(),
		})
	}
	return v1.NewV1DomainsShorternedOK().WithPayload(&models.V1DomainsShorterned{
		Domains: converters.ToDomainCounterList(domainCounter),
	})
}

func (v *V1Handler) Redirect(params v1.V1TinyurlRedirectParams) middleware.Responder {
	tinyURL, err := v.TinyUrlService.GetOriginalURL(params.Tinyurl)
	if err != nil {
		return v1.NewV1TinyurlRedirectInternalServerError().WithPayload(&models.V1InternalServerError{
			Error: err.Error(),
		})
	}
	return v1.NewV1TinyurlRedirectFound().WithLocation(tinyURL.URL)
}
