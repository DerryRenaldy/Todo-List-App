package dtotodo

type CreateTodoRequest struct {
	Title           string `json:"title"`
	ActivityGroupId int    `json:"activity_group_id"`
	Priority        string `json:"priority"`
	IsActive        bool   `json:"is_active"`
}
