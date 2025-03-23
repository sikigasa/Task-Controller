package usecase

import (
	task "github.com/sikigasa/task-controller/proto/v1"
)

type taskService struct {
	task.UnimplementedTaskServiceServer
}
