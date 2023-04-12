package dtotodo

type UpdateTodoRequest struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
	Status   string `json:"status"`
}
