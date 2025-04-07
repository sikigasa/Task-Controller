package domain

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	IsEnd bool `json:"is_end"`

	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
	LimitedAt string `json:"limited_at"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
