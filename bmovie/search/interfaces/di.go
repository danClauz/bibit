package interfaces

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(New); err != nil {
		return errors.Wrap(err, "failed to provide app service")
	}
	return nil
}
