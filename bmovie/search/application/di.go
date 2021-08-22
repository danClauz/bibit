package application

import (
	"github.com/danClauz/bibit/bmovie/search/external/omdb"
	"github.com/danClauz/bibit/bmovie/search/repository"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(New); err != nil {
		return errors.Wrap(err, "failed to provide application service")
	}
	if err := container.Provide(repository.New); err != nil {
		return errors.Wrap(err, "failed to provide repository service")
	}
	if err := container.Provide(omdb.NewClient); err != nil {
		return errors.Wrap(err, "failed to provide omdb external client")
	}
	return nil
}
