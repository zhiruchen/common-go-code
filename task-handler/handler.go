package taskhandler

import (
	"context"
)

type Task interface {
	Name() string
	Execute(ctx context.Context) error
	CanExecute() bool
}

type Handler interface {
	Handle(ctx context.Context) error
	SetRunTime(...Task) error
	Stop() error
}
