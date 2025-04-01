package domain

import "time"

type CreateTaskParam struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	LimitedAt   time.Time `json:"limited_at"`
	IsEnd       bool      `json:"is_end"`
}

type GetTaskParam struct {
	ID string `json:"id"`
}

type ListTaskParam struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type UpdateTaskParam struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsEnd       bool   `json:"is_end"`
}

type DeleteTaskParam struct {
	ID string `json:"id"`
}
