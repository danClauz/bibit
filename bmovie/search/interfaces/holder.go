package interfaces

import "go.uber.org/dig"

type Holder struct {
	dig.In
	Interfaces
}