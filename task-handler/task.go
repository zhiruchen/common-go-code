package taskhandler

import "context"

type TaskImpl struct {
	
}

func (t TaskImpl) Name() string {
	panic("implement me")
}

func (t TaskImpl) Execute(ctx context.Context) error {
	panic("implement me")
}

func (t TaskImpl) CanExecute() bool {
	panic("implement me")
}

func NewTask() Task {
	return &TaskImpl{}
}
