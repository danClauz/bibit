package application

import "go.uber.org/dig"

type Holder struct {
	dig.In
	Application
}
