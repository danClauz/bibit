package infrastructure

import (
	"github.com/danClauz/bibit/bmovie/search/infrastructure/gateway"
	"github.com/danClauz/bibit/bmovie/search/infrastructure/server"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(server.New); err != nil {
		return errors.Wrap(err, "failed to provide http controller")
	}
	if err := container.Provide(gateway.New); err != nil {
		return errors.Wrap(err, "failed to provide gateway controller")
	}
	return nil
}
