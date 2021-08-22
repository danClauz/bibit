package repository

import "go.uber.org/dig"

type Holder struct {
	dig.In
	Repository
}
