package shared

import (
	"github.com/danClauz/bibit/bmovie/search/shared/config"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"io"
)

type Holder struct {
	dig.In
	Config *config.EnvConfig
	Logger *logrus.Logger
	Echo   *echo.Echo
	Mysql  *gorm.DB
}

func (h *Holder) Close() {
	h.Logger.Info("closing resource")

	var i interface{} = nil

	i = h.Logger
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}
	i = h.Echo
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}
	i = h.Mysql
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}
}
