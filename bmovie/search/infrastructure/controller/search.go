package controller

import (
	"errors"
	"strconv"

	"github.com/danClauz/bibit/bmovie/search/shared/logtag"
	"github.com/danClauz/bibit/bmovie/search/shared/response"
	"github.com/danClauz/bibit/bmovie/search/shared/utils"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) SearchMovie(ctx echo.Context) error {
	reqId := utils.FetchRequestIdFromContext(ctx)

	search := ctx.QueryParam("s")
	page := ctx.QueryParam("p")
	id := ctx.QueryParam("i")

	logger := ctrl.sh.Logger.WithFields(map[string]interface{}{})

	logger.Infof(logtag.RequestTmpl, logtag.PerformSearchMovie)

	var resp interface{}
	var respErr error = nil

	if search != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			p = 1
		}

		resp, respErr = ctrl.ih.SearchMovie(ctx.Request().Context(), reqId, search, p)
	} else if id != "" {
		resp, respErr = ctrl.ih.DetailMovie(ctx.Request().Context(), reqId, id)
	} else {
		err := response.Error(response.BAD_REQUEST, errors.New("invalid query input"))
		return response.Body(ctx, nil, err)
	}

	if respErr != nil {
		logger.Errorf(logtag.ErrorTmpl, logtag.PerformSearchMovie, respErr)
		return response.Body(ctx, nil, respErr)
	}

	return response.Body(ctx, resp, nil)
}
