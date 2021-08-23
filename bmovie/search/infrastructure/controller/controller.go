package controller

import (
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
)

type Controller struct {
	sh shared.Holder
	ih interfaces.Holder
}

func New(sh shared.Holder, ih interfaces.Holder) *Controller {
	return &Controller{
		sh: sh,
		ih: ih,
	}
}
