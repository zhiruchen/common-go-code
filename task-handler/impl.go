package taskhandler

import (
	"context"
)

type DefaultTaskHandler struct{}

func (d DefaultTaskHandler) Handle(ctx context.Context) error {
	panic("implement me")
}

func (d DefaultTaskHandler) SetRunTime(...Task) error {
	panic("implement me")
}

func (d DefaultTaskHandler) Stop() error {
	panic("implement me")
}

func NewTaskHandler() Handler {
	return &DefaultTaskHandler{}
}
