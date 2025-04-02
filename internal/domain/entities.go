package domain

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	IsEnd bool `json:"is_end"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
