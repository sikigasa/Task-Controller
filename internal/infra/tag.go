package infra

import (
	"context"

	"github.com/sikigasa/task-controller/internal/domain"
)

func (t *taskRepo) CreateTag(ctx context.Context, arg domain.CreateTaskParam) error {
	const query = `INSERT INTO Tag (id, name) VALUES ($1,$2)`

	row := t.db.QueryRowContext(ctx, query, arg.ID, arg.Title, arg.Description, arg.IsEnd)

	return row.Err()
}

func (t *taskRepo) ListTag(ctx context.Context, arg domain.ListTagParam) ([]domain.Tag, error) {
	const query = `SELECT id, name FROM Tag LIMIT $1 OFFSET $2`

	rows, err := t.db.QueryContext(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []domain.Tag
	for rows.Next() {
		var tag domain.Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (t *taskRepo) DeleteTag(ctx context.Context, arg domain.DeleteTagParam) error {
	const query = `DELETE FROM Tag WHERE id = $1`

	row := t.db.QueryRowContext(ctx, query, arg.ID)

	return row.Err()
}
