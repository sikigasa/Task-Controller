package usecase

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sikigasa/task-controller/internal/domain"
	"github.com/sikigasa/task-controller/internal/infra"
	postgres "github.com/sikigasa/task-controller/internal/infra/driver"
	task "github.com/sikigasa/task-controller/proto/v1"
)

type taskService struct {
	task.UnimplementedTaskServiceServer
	taskRepo    infra.TaskRepo
	taskTagRepo infra.TaskTagRepo
	tx          postgres.Transaction
}

func NewTaskService(taskRepo infra.TaskRepo, taskTagRepo infra.TaskTagRepo, tx postgres.Transaction) task.TaskServiceServer {
	return &taskService{
		taskRepo:    taskRepo,
		taskTagRepo: taskTagRepo,
		tx:          tx,
	}
}

func (t *taskService) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.CreateTaskResponse, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	err = t.tx.WithTransaction(ctx, func(tx *sql.Tx) error {
		param := domain.CreateTaskParam{
			ID:          uuid.String(),
			Title:       req.Title,
			Description: req.Description,
			IsEnd:       false,
		}

		if err := t.taskRepo.CreateTask(ctx, param); err != nil {
			return err
		}
		for _, tagID := range req.TagIds {
			taskTagParam := domain.CreateTaskTagParam{
				TaskID: param.ID,
				TagID:  tagID,
			}
			if err := t.taskTagRepo.CreateTaskTag(ctx, taskTagParam); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
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

	taskDetail, err := t.taskRepo.GetTask(ctx, param)
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

	tasks, err := t.taskRepo.ListTask(ctx, param)
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

func (t *taskService) UpdateTask(ctx context.Context, req *task.UpdateTaskRequest) (*task.UpdateTaskResponse, error) {
	err := t.tx.WithTransaction(ctx, func(tx *sql.Tx) error {
		param := domain.UpdateTaskParam{
			ID:          req.Id,
			Title:       req.Title,
			Description: req.Description,
		}
		if err := t.taskRepo.UpdateTask(ctx, param); err != nil {
			return err
		}
		if err := t.taskTagRepo.DeleteTaskTags(ctx, domain.DeleteTaskTagParam{TaskID: req.Id}); err != nil {
			return err
		}
		for _, tagID := range req.TagIds {
			taskTagParam := domain.CreateTaskTagParam{
				TaskID: param.ID,
				TagID:  tagID,
			}
			if err := t.taskTagRepo.CreateTaskTag(ctx, taskTagParam); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &task.UpdateTaskResponse{}, nil
}

func (t *taskService) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.DeleteTaskResponse, error) {
	err := t.tx.WithTransaction(ctx, func(tx *sql.Tx) error {
		if err := t.taskTagRepo.DeleteTaskTags(ctx, domain.DeleteTaskTagParam{TaskID: req.Id}); err != nil {
			return err
		}

		param := domain.DeleteTaskParam{
			ID: req.Id,
		}
		if err := t.taskRepo.DeleteTask(ctx, param); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &task.DeleteTaskResponse{}, nil
}
