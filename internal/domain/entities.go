package domain

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsEnd       bool   `json:"is_end"`
}
