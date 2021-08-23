package omdb

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/danClauz/bibit/bmovie/search/shared"
)

//go:generate mockgen -source=client.go -destination=mocks/omdb_mock.go -package=omdb_mock
type (
	Client interface {
		SearchMovie(ctx context.Context, reqId, searchKey string, page int) (*SearchResponse, error)
		SearchMovieByImdbId(ctx context.Context, reqId, searchKey string) (*SearchByTitleResponse, error)
	}
	client struct {
		sh   shared.Holder
		c    *http.Client
		host string
		key  string
	}
)

func NewClient(sh shared.Holder) Client {
	return &client{
		sh: sh,
		c: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 50,
				DialContext: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 10 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 5 * time.Second,
			},
			Timeout: 5 * time.Second,
		},
		host: sh.Config.Omdb.Host,
		key:  sh.Config.Omdb.Key,
	}
}
