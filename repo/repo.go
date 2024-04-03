package repo

import (
	"github.com/pkg/errors"

	"github.com/LochanRn/tiny-url-server/domain"

	local_store "github.com/LochanRn/tiny-url-server/repo/local_store"
	"github.com/LochanRn/tiny-url-server/repo/models"
)

type Repo interface {
	CreateURL(url *models.URL) (string, bool, error)
	GetURL(url string) (*models.URL, error)
	IncrementDomainCounter(domain string) error
	GetTop3DomainCount() (models.CounterList, error)
}

func NewRepo(dbType domain.DBType) (Repo, error) {
	switch dbType {
	case domain.LocalStore:
		return local_store.NewLocalStore()
	default:
		return nil, errors.Errorf("invalid db type %v", dbType)
	}
}
