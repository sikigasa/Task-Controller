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
		Id: uuid.String(),
	}, nil
}

func (t *taskService) GetTask(ctx context.Context, req *task.GetTaskRequest) (*task.GetTaskResponse, error) {
	param := domain.GetTaskParam{
		ID: req.Id,
	}

	taskDetail, err := t.taskRepo.Get(ctx, param)
	if err != nil {
		return nil, err
	}

	return &task.GetTaskResponse{
		Task: &task.Task{
			Id:          taskDetail.ID,
			Title:       taskDetail.Title,
			Description: taskDetail.Description,
			IsEnd:       taskDetail.IsEnd,
		},
	}, nil
}

func (t *taskService) ListTask(ctx context.Context, req *task.ListTaskRequest) (*task.ListTaskResponse, error) {
	param := domain.ListTaskParam{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	tasks, err := t.taskRepo.GetAll(ctx, param)
	if err != nil {
		return nil, err
	}

	var taskList []*task.Task
	for _, taskDetail := range tasks {
		taskList = append(taskList, &task.Task{
			Id:          taskDetail.ID,
			Title:       taskDetail.Title,
			Description: taskDetail.Description,
			IsEnd:       taskDetail.IsEnd,
		})
	}

	return &task.ListTaskResponse{
		Tasks: taskList,
	}, nil
}
