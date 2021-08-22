package di

import (
	"github.com/danClauz/bibit/bmovie/search/shared/config"
	"github.com/danClauz/bibit/common/go/logs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

func NewLogger(cfg *config.EnvConfig) (*logrus.Logger, error) {
	return logs.New(&logs.Option{
		FilePath:     cfg.Logger.FilePath,
		FileName:     cfg.Logger.FileName,
		Formatter:    logs.Formatter(strings.ToUpper(cfg.Logger.Formatter)),
		Stdout:       cfg.Logger.Stdout,
		ReportCaller: cfg.Logger.ReportCaller,
	})
}

func init() {
	if err := Container.Provide(NewLogger); err != nil {
		panic(errors.Wrap(err, "failed to provide logger"))
	}
}
