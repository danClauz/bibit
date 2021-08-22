package di

import (
	"github.com/danClauz/bibit/bmovie/search/application"
	"github.com/danClauz/bibit/bmovie/search/infrastructure"
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register shared container"))
	}
	if err := infrastructure.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register infrastructure container"))
	}
	if err := interfaces.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register interfaces container"))
	}
	if err := application.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register application container"))
	}
}
