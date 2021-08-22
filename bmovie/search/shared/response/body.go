package response

import (
	"github.com/danClauz/bibit/bmovie/search/shared/dto"
	"github.com/labstack/echo/v4"
)

type (
	StdResp interface {
		GetHttpStatus() int
		GetCode() string
		GetMessage() string
	}
)

func Body(ctx echo.Context, data interface{}, err interface{}) error {
	if err == nil {
		resp := SUCCESS
		return ctx.JSON(resp.GetHttpStatus(), &dto.HttpRespBody{
			Code:    resp.GetCode(),
			Message: resp.GetMessage(),
			Data:    data,
		})
	}

	var resp StdError

	switch e := err.(type) {
	case StdError:
		resp = e
	case error:
		resp = Error(SYSTEM_ERROR, e)
	}

	return ctx.JSON(resp.GetHttpStatus(), &dto.HttpRespBody{
		Code:    resp.GetCode(),
		Message: resp.GetMessage(),
		Data:    resp.GetErrors(),
	})
}
