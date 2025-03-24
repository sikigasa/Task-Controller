package infra

import (
	"context"
	"database/sql"

	"github.com/sikigasa/task-controller/internal/domain"
)

type taskRepo struct {
	db *sql.DB
}

type TaskRepo interface {
	Create(ctx context.Context, arg domain.CreateTaskParam) error
	Get(ctx context.Context, arg domain.GetTaskParam) (*domain.Task, error)

	Delete(ctx context.Context, arg domain.DeleteTaskParam) error
}

func NewTaskRepo(db *sql.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (t *taskRepo) Create(ctx context.Context, arg domain.CreateTaskParam) error {
	const query = `INSERT INTO task (title, description, authority) VALUES ($1,$2,$3)`

	row := t.db.QueryRowContext(ctx, query, arg.Title, arg.Description, arg.IsEnd)

	return row.Err()
}

func (t *taskRepo) Get(ctx context.Context, arg domain.GetTaskParam) (*domain.Task, error) {
	const query = `SELECT task_id,project_id,authority FROM tasks WHERE task_id = $1`
	row := t.db.QueryRowContext(ctx, query, arg.ID)
	var task domain.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description); err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepo) Delete(ctx context.Context, arg domain.DeleteTaskParam) error {
	const query = `DELETE FROM tasks WHERE project_id = $1`
	row, err := t.db.ExecContext(ctx, query, arg.ID)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
