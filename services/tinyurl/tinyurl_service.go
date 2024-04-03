package tinyurl

import (
	"github.com/LochanRn/tiny-url-server/domain"
	"github.com/LochanRn/tiny-url-server/repo"
	"github.com/pkg/errors"
)

type TinyURLService struct {
	repo repo.Repo
}

func NewTinyURLService(dbType domain.DBType) (*TinyURLService, error) {
	repo, err := repo.NewRepo(dbType)
	if err != nil {
		return nil, errors.Wrap(err, "error while creating new repo")
	}
	return &TinyURLService{
		repo: repo,
	}, nil
}
