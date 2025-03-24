package usecase

import (
	"context"

	"github.com/sikigasa/task-controller/internal/infra"
	task "github.com/sikigasa/task-controller/proto/v1"
)

type taskService struct {
	task.UnimplementedTaskServiceServer
	taskRepo infra.TaskRepo
}

func NewTaskService(taskRepo infra.TaskRepo) task.TaskServiceServer {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (t *taskService) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.CreateTaskResponse, error) {
	param := task.CreateTaskParam{
		Title:       req.Title,
		Description: req.Description,
		IsEnd:       req.IsEnd,
	}

	return &task.CreateTaskResponse{
		Id: "generated-task-id", // Replace with actual task ID
	}, nil
}
