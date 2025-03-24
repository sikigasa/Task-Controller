package domain

type CreateTaskParam struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsEnd       bool   `json:"is_end"`
}

type GetTaskParam struct {
	ID string `json:"id"`
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
