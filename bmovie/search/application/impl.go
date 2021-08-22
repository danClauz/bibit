package application

import (
	"context"
	"github.com/danClauz/bibit/bmovie/search/external/omdb"
	"github.com/danClauz/bibit/bmovie/search/model"
	"github.com/danClauz/bibit/bmovie/search/repository"
	"github.com/danClauz/bibit/bmovie/search/shared"
)

//go:generate mockgen -source=impl.go -destination=mocks/impl_mock.go -package=application_mock Application
type (
	Application interface {
		ListMovieByTitle(ctx context.Context, reqId, searchKey string, page int) (*model.SearchHistory, error)
		MovieDetailByImdbId(ctx context.Context, reqId, searchKey string) (*model.SearchHistory, error)
	}
	app struct {
		rh         repository.Holder
		sh         shared.Holder
		omdbClient omdb.Client
	}
)

func New(rh repository.Holder, sh shared.Holder, omdbClient omdb.Client) Application {
	return &app{
		rh:         rh,
		sh:         sh,
		omdbClient: omdbClient,
	}
}

func (svc *app) ListMovieByTitle(ctx context.Context, reqId, searchKey string, page int) (*model.SearchHistory, error) {
	resp, err := svc.omdbClient.SearchMovie(ctx, reqId, searchKey, page)
	if err != nil {
		return nil, err
	}

	record := &model.SearchHistory{
		RequestID: reqId,
		SearchKey: searchKey,
		Page:      page,
		Result:    resp,
	}

	if err := svc.rh.Store(ctx, reqId, record); err != nil {
		return nil, err
	}

	return record, nil
}

func (svc *app) MovieDetailByImdbId(ctx context.Context, reqId, searchKey string) (*model.SearchHistory, error) {
	resp, err := svc.omdbClient.SearchMovieByImdbId(ctx, reqId, searchKey)
	if err != nil {
		return nil, err
	}

	record := &model.SearchHistory{
		RequestID: reqId,
		SearchKey: searchKey,
		Result:    resp,
	}

	if err := svc.rh.Store(ctx, reqId, record); err != nil {
		return nil, err
	}

	return record, nil
}
