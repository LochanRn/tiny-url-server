package converters

import (
	"github.com/LochanRn/tiny-url-api/gen/models"
	"github.com/LochanRn/tiny-url-server/domain"
)

func ToV1TinyURL(tinyURL domain.TinyURL) *models.V1TinyurlRedirectFoundBody {
	return &models.V1TinyurlRedirectFoundBody{
		URL:     tinyURL.URL,
		Tinyurl: tinyURL.TinyURL,
	}
}
