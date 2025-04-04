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
	CreateTask(ctx context.Context, arg domain.CreateTaskParam) error
	GetTask(ctx context.Context, arg domain.GetTaskParam) (*domain.Task, error)
	GetAllTask(ctx context.Context, arg domain.ListTaskParam) ([]domain.Task, error)
	UpdateTask(ctx context.Context, arg domain.UpdateTaskParam) error
	DeleteTask(ctx context.Context, arg domain.DeleteTaskParam) error
}

func NewTaskRepo(db *sql.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (t *taskRepo) CreateTask(ctx context.Context, arg domain.CreateTaskParam) error {
	const query = `INSERT INTO task (id, title, description, limited_at, is_end) VALUES ($1,$2,$3)`

	row := t.db.QueryRowContext(ctx, query, arg.ID, arg.Title, arg.Description, arg.LimitedAt, arg.IsEnd)

	return row.Err()
}

func (t *taskRepo) GetTask(ctx context.Context, arg domain.GetTaskParam) (*domain.Task, error) {
	const query = `SELECT task_id,project_id,authority FROM task WHERE task_id = $1`
	row := t.db.QueryRowContext(ctx, query, arg.ID)
	var task domain.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description); err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepo) GetAllTask(ctx context.Context, arg domain.ListTaskParam) ([]domain.Task, error) {
	const query = `SELECT task_id,project_id,authority FROM task LIMIT $1 OFFSET $2`
	rows, err := t.db.QueryContext(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil

}

func (t *taskRepo) UpdateTask(ctx context.Context, arg domain.UpdateTaskParam) error {
	const query = `UPDATE task SET title = $1, description = $2 WHERE task_id = $3`
	row := t.db.QueryRowContext(ctx, query, arg.Title, arg.Description, arg.ID)
	return row.Err()
}

func (t *taskRepo) DeleteTask(ctx context.Context, arg domain.DeleteTaskParam) error {
	const query = `DELETE FROM task WHERE project_id = $1`
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
