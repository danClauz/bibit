package shared

import (
	"github.com/danClauz/bibit/bmovie/search/shared/config"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(config.NewConfiguration); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
