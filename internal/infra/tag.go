package infra

import (
	"context"

	"github.com/sikigasa/task-controller/internal/domain"
)

func (t *taskRepo) Create(ctx context.Context, arg domain.CreateTaskParam) error {
	const query = `INSERT INTO task (id, title, description, limited_at, is_end) VALUES ($1,$2,$3)`

	row := t.db.QueryRowContext(ctx, query, arg.ID, arg.Title, arg.Description, arg.IsEnd)

	return row.Err()
}
