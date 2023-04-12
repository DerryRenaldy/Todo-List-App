package todostore

const (
	QueryCreateTodo        = `INSERT INTO todoApp.todos (activity_group_id, title, priority, is_active) VALUES (?, ?, ?, ?);`
	QueryGetOneTodoById    = `SELECT todo_id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todoApp.todos WHERE todo_id=?;`
	QueryGetTodoList       = `SELECT todo_id, activity_group_id, title, is_active, priority, created_at, updated_at FROM todoApp.todos`
	QueryUpdateDeletedTime = `UPDATE todoApp.todos SET deleted_at=? WHERE todo_id=?;`
)
