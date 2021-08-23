package controller

import (
	"github.com/danClauz/bibit/bmovie/search/shared/response"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) HealthCheck(ctx echo.Context) error {
	return response.Body(ctx, nil, nil)
}
