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
	GetTaskTags(ctx context.Context, arg domain.GetTaskTagParam) ([]domain.Tag, error)
	DeleteTaskTags(ctx context.Context, arg domain.DeleteTaskTagParam) error
}

func NewTaskTagRepo(db *sql.DB) TaskTagRepo {
	return &taskTagRepo{db: db}
}

func (t *taskTagRepo) CreateTaskTag(ctx context.Context, arg domain.CreateTaskTagParam) error {
	const query = `INSERT INTO task_tag (task_id, tag_id) VALUES ($1,$2)`

	row := t.db.QueryRowContext(ctx, query, arg.TaskID, arg.TagID)

	return row.Err()
}

func (t *taskTagRepo) GetTaskTags(ctx context.Context, arg domain.GetTaskTagParam) ([]domain.Tag, error) {
	const query = `SELECT * FROM task_tag WHERE task_id = $1`

	rows, err := t.db.QueryContext(ctx, query, arg.TaskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taskTags []domain.Tag
	for rows.Next() {
		var taskTag domain.Tag
		if err := rows.Scan(&taskTag.ID, &taskTag.Name); err != nil {
			return nil, err
		}
		taskTags = append(taskTags, taskTag)
	}
	return taskTags, nil
}

func (t *taskTagRepo) DeleteTaskTags(ctx context.Context, arg domain.DeleteTaskTagParam) error {
	const query = `DELETE FROM task_tag WHERE task_id = $1`
	row := t.db.QueryRowContext(ctx, query, arg.TaskID)

	return row.Err()
}
