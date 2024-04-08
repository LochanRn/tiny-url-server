package converters

import (
	"github.com/LochanRn/tiny-url-api/gen/models"
	"github.com/LochanRn/tiny-url-server/domain"
)

func ToV1TinyURL(tinyURL domain.TinyURL) *models.V1Tinyurl {
	return &models.V1Tinyurl{
		URL: tinyURL.URL,
	}
}

func ToDomainCounterList(counters domain.CounterList) []*models.V1Domain {
	var domainCounters []*models.V1Domain
	for _, counter := range counters {
		domainCounters = append(domainCounters, ToV1DomainCounter(counter))
	}
	return domainCounters
}

func ToV1DomainCounter(counter domain.Counter) *models.V1Domain {
	return &models.V1Domain{
		Domains: counter.Domain,
		Count:   int64(counter.Count),
	}
}
