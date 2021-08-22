package di

import (
	ech "github.com/danClauz/bibit/common/go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewEcho(logger *ech.LoggerWrapper) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Logger = logger
	return e, nil
}

func NewLoggerWrapper(logger *logrus.Logger) *ech.LoggerWrapper {
	return ech.NewLoggerWrapper(logger)
}

func init() {
	if err := Container.Provide(NewEcho); err != nil {
		panic(errors.Wrap(err, "failed to provide echo"))
	}
	if err := Container.Provide(NewLoggerWrapper); err != nil {
		panic(errors.Wrap(err, "failed to provide echo logger wrapper"))
	}
}
