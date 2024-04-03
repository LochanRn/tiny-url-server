package tinyurl

import (
	"net/url"
	"strings"
	"time"

	"github.com/LochanRn/tiny-url-server/repo/models"
	"github.com/LochanRn/tiny-url-server/utils/base62"
	"github.com/pkg/errors"
)

func (t *TinyURLService) CreateTinyURL(url string) (string, error) {
	tinyURL := base62.Base62Encode(url)
	u := &models.URL{
		OriginalURL: url,
		TinyURL:     tinyURL,
		CreatedAt:   time.Now(),
	}

	tinyURL, ok, err := t.repo.CreateURL(u)
	if err != nil {
		return "", errors.Wrapf(err, "error while creating new tiny url %s", url)
	}

	if ok {
		err = t.incrementDomainCounter(url)
		if err != nil {
			return tinyURL, errors.Wrapf(err, "error while incrementing domain counter %s", url)
		}
	}
	return tinyURL, nil
}

func (t *TinyURLService) GetOriginalURL(tinyURL string) (string, error) {
	return base62.Base62Decode(tinyURL), nil
}

func (t *TinyURLService) incrementDomainCounter(url string) error {
	domain, err := extractDomain(url)
	if err != nil {
		return errors.Wrapf(err, "error while extracting domain from %s", url)
	}
	err = t.repo.IncrementDomainCounter(domain)
	if err != nil {
		return errors.Wrapf(err, "error while incrementing domain counter %s", url)
	}
	return nil
}

func extractDomain(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	host := parsedURL.Hostname()
	return strings.TrimPrefix(host, "www."), nil
}
