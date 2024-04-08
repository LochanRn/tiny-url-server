package converter

import (
	"github.com/LochanRn/tiny-url-server/domain"
	"github.com/LochanRn/tiny-url-server/repo/models"
)

func ToDomainCounterList(cl models.CounterList) domain.CounterList {
	var dcl domain.CounterList
	for _, c := range cl {
		dcl = append(dcl, domain.Counter{
			Domain: c.Domain,
			Count:  c.Count,
		})
	}
	return dcl
}

func ToDomainTinyURL(tu *models.URL) *domain.TinyURL {
	return &domain.TinyURL{
		URL:     tu.OriginalURL,
		TinyURL: tu.TinyURL,
	}
}
