package task

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "task"
)

var (
	validate = validator.New()
)

func NewCreateTaskRequst() *CreateTaskRequst {
	return &CreateTaskRequst{
		Params: map[string]string{},
	}
}

func (req *CreateTaskRequst) Validate() error {
	return validate.Struct(req)
}

func CreateTask(req *CreateTaskRequst) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ins := NewDefaultTask()
	ins.Data = req

	return ins, nil
}

func NewDefaultTask() *Task {
	return &Task{
		Data:   &CreateTaskRequst{},
		Status: &Status{},
	}
}

func (s *Task) Run() {
	s.Status.StartAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_RUNNING
}

func (s *Task) Failed(message string) {
	s.Status.EndAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_FAILED
	s.Status.Message = message
}

func (s *Task) Success() {
	s.Status.EndAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_SUCCESS
}
