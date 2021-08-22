package omdb

import (
	"context"
	"fmt"
	"github.com/danClauz/bibit/bmovie/search/shared/errors"
	"github.com/danClauz/bibit/bmovie/search/shared/logtag"
	"github.com/danClauz/bibit/bmovie/search/shared/utils"
	"net/http"
)

const (
	searchMovieEndpoint         = "/?apikey=%s&s=%s&page=%d"
	searchMovieByImdbIdEndpoint = "/?apikey=%s&i=%s"
)

func (c *client) SearchMovie(ctx context.Context, reqId, searchKey string, page int) (*SearchResponse, error) {
	logger := c.sh.Logger.WithFields(map[string]interface{}{
		utils.XRequestId: reqId,
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.host+fmt.Sprintf(searchMovieEndpoint, c.key, searchKey, page), nil)
	if err != nil {
		logger.Errorf(logtag.ErrorTmpl, logtag.ErrBuildHttpRequest, err)
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		logger.Errorf(logtag.ErrorTmpl, logtag.ErrDoHTTPCall, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		r := new(SearchResponse)
		if err := utils.ReadAll(resp.Body, r); err != nil {
			logger.Errorf(logtag.ErrorTmpl, logtag.ErrReadResponseBody, err)
			return nil, err
		}

		return r, nil
	}

	err = errors.ErrUnhandledHttpStatus(resp.StatusCode)
	logger.Errorf(logtag.ErrorTmpl, logtag.ErrUnhandledHTTPStatus, err)
	return nil, err
}

func (c *client) SearchMovieByImdbId(ctx context.Context, reqId, searchKey string) (*SearchByTitleResponse, error) {
	logger := c.sh.Logger.WithFields(map[string]interface{}{
		utils.XRequestId: reqId,
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.host+fmt.Sprintf(searchMovieByImdbIdEndpoint, c.key, searchKey), nil)
	if err != nil {
		logger.Errorf(logtag.ErrorTmpl, logtag.ErrBuildHttpRequest, err)
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		logger.Errorf(logtag.ErrorTmpl, logtag.ErrDoHTTPCall, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		r := new(SearchByTitleResponse)
		if err := utils.ReadAll(resp.Body, r); err != nil {
			logger.Errorf(logtag.ErrorTmpl, logtag.ErrReadResponseBody, err)
			return nil, err
		}

		return r, nil
	}

	err = errors.ErrUnhandledHttpStatus(resp.StatusCode)
	logger.Errorf(logtag.ErrorTmpl, logtag.ErrUnhandledHTTPStatus, err)
	return nil, err
}
