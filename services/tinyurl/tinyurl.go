package tinyurl

import (
	"time"

	"github.com/LochanRn/tiny-url-server/domain"
	"github.com/LochanRn/tiny-url-server/repo/models"
	"github.com/LochanRn/tiny-url-server/services/tinyurl/converter"
	"github.com/LochanRn/tiny-url-server/utils/random"
	urlutil "github.com/LochanRn/tiny-url-server/utils/url"
	"github.com/pkg/errors"
)

func (t *TinyURLService) CreateTinyURL(url string) (*domain.TinyURL, error) {
	u := &models.URL{
		OriginalURL: url,
		TinyURL:     random.RandomNumberInRange(),
		CreatedAt:   time.Now(),
	}

	tinyURL, ok, err := t.repo.CreateURL(u)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating new tiny url %s", url)
	}

	tu := &domain.TinyURL{
		URL:               u.OriginalURL,
		TinyURL:           domain.TinyUrlPrefix + tinyURL,
		CreationTimeStamp: domain.Time(u.CreatedAt),
	}

	if ok {
		if err := t.IncrementDomainCounter(url); err != nil {
			return tu, errors.Wrapf(err, "error while incrementing domain counter %s", url)
		}
	}
	return tu, nil
}

func (t *TinyURLService) GetOriginalURL(tinyURL string) (*domain.TinyURL, error) {
	url, err := t.repo.GetURL(tinyURL)
	if err != nil {
		return nil, errors.Wrapf(err, "error while fetching original url for %s", tinyURL)
	}
	url.TinyURL = domain.TinyUrlPrefix + url.TinyURL
	return converter.ToDomainTinyURL(url), nil
}

func (t *TinyURLService) IncrementDomainCounter(url string) error {
	domain := urlutil.GetDomainFromURL(url)
	err := t.repo.IncrementDomainCounter(domain)
	if err != nil {
		return errors.Wrapf(err, "error while incrementing domain counter %s", url)
	}
	return nil
}

func (t *TinyURLService) GetDomainsShortened() (domain.CounterList, error) {
	counters, err := t.repo.GetDomainsShortened()
	if err != nil {
		return nil, errors.Wrap(err, "error while fetching top 3 domain counter")
	}
	return converter.ToDomainCounterList(counters), nil
}
