package infra

import (
	"context"
	"database/sql"

	"github.com/sikigasa/task-controller/internal/domain"
)

type taskTagRepo struct {
	db *sql.DB
}

type TaskTagRepo interface {
	CreateTaskTag(ctx context.Context, arg domain.CreateTaskTagParam) error
	DeleteTaskTag(ctx context.Context, arg domain.DeleteTaskTagParam) error
}

func NewTaskTagRepo(db *sql.DB) TaskTagRepo {
	return &taskTagRepo{db: db}
}

func (t *taskTagRepo) CreateTaskTag(ctx context.Context, arg domain.CreateTaskTagParam) error {
	const query = `INSERT INTO task_tag (task_id, tag_id) VALUES ($1,$2)`

	row := t.db.QueryRowContext(ctx, query, arg.TaskID, arg.TagID)

	return row.Err()
}

func (t *taskTagRepo) DeleteTaskTag(ctx context.Context, arg domain.DeleteTaskTagParam) error {
	const query = `DELETE FROM task_tag WHERE task_id = $1 AND tag_id = $2`

	row := t.db.QueryRowContext(ctx, query, arg.TaskID, arg.TagID)

	return row.Err()
}
