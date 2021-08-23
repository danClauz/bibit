package interfaces

import (
	"context"

	"github.com/danClauz/bibit/bmovie/search/application"
	"github.com/danClauz/bibit/bmovie/search/model"
	"github.com/danClauz/bibit/bmovie/search/shared"
)

//go:generate mockgen -source=impl.go -destination=mocks/impl_mock.go -package=interfaces_mock Interfaces
type (
	Interfaces interface {
		SearchMovie(ctx context.Context, reqId, searchKey string, page int) (*model.SearchHistory, error)
		DetailMovie(ctx context.Context, reqId, searchKey string) (*model.SearchHistory, error)
	}
	interfaces struct {
		sh shared.Holder
		ah application.Holder
	}
)

func New(sh shared.Holder, ah application.Holder) Interfaces {
	return &interfaces{
		sh: sh,
		ah: ah,
	}
}

func (i *interfaces) SearchMovie(ctx context.Context, reqId, searchKey string, page int) (*model.SearchHistory, error) {
	return i.ah.ListMovieByTitle(ctx, reqId, searchKey, page)
}

func (i *interfaces) DetailMovie(ctx context.Context, reqId, searchKey string) (*model.SearchHistory, error) {
	return i.ah.MovieDetailByImdbId(ctx, reqId, searchKey)
}
