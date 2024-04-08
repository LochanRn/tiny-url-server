package localstore

import (
	"errors"
	"sort"

	"github.com/LochanRn/tiny-url-server/repo/models"
)

type LocalStore struct {
	store map[string]interface{}
}

const (
	urls    = "urls"
	counter = "counter"
)

func NewLocalStore() (*LocalStore, error) {
	store := make(map[string]interface{})
	store[urls] = make(map[string]*models.URL)
	store[counter] = make(map[string]*models.Counter)
	return &LocalStore{store: store}, nil
}

func (lc *LocalStore) CreateURL(url *models.URL) (string, bool, error) {
	u, ok := lc.store[urls].(map[string]*models.URL)
	if !ok {
		return "", false, errors.New("urls object not found")
	}
	if val, ok := lc.findURL(url.OriginalURL); ok {
		return val.TinyURL, false, nil
	}
	u[url.TinyURL] = url
	lc.store[urls] = u
	return url.TinyURL, true, nil
}

func (lc *LocalStore) findURL(originalURL string) (*models.URL, bool) {
	urls, ok := lc.store[urls].(map[string]*models.URL)
	if !ok {
		return nil, false
	}
	for _, v := range urls {
		if v.OriginalURL == originalURL {
			return v, true
		}
	}
	return nil, false
}

func (lc *LocalStore) GetURL(url string) (*models.URL, error) {
	urls, ok := lc.store[urls].(map[string]*models.URL)
	if !ok {
		return nil, errors.New("urls object not found")
	}
	val, ok := urls[url]
	if !ok {
		return nil, errors.New("tiny url not found")
	}
	return val, nil
}

func (lc *LocalStore) IncrementDomainCounter(domain string) error {
	c, ok := lc.store[counter].(map[string]*models.Counter)
	if !ok {
		return errors.New("counter object not found")
	}
	if val, ok := c[domain]; ok {
		val.Count++
	} else {
		c[domain] = &models.Counter{Count: 1, Domain: domain}
	}
	lc.store[counter] = c
	return nil
}

func (lc *LocalStore) GetDomainsShortened() (models.CounterList, error) {
	counter, ok := lc.store[counter].(map[string]*models.Counter)
	if !ok {
		return nil, errors.New("counter object not found")
	}

	list := make(models.CounterList, 0, len(counter))
	for _, v := range counter {
		list = append(list, *v)
	}
	sort.Sort(list)

	if len(list) > 3 {
		list = list[:3]
	}
	return list, nil
}
