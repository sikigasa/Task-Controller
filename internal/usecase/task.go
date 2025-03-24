package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sikigasa/task-controller/internal/domain"
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
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	param := domain.CreateTaskParam{
		ID:          uuid.String(),
		Title:       req.Title,
		Description: req.Description,
		IsEnd:       false,
	}

	if err := t.taskRepo.Create(ctx, param); err != nil {
		return nil, err
	}

	return &task.CreateTaskResponse{
		Id: r, // Replace with actual task ID
	}, nil
}
