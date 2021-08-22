package utils

import "github.com/labstack/echo/v4"

const (
	XRequestId = "x-request-id"
)

func FetchRequestIdFromContext(ctx echo.Context) string {
	return ctx.Request().Header.Get(XRequestId)
}