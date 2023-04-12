package todostore

import (
	"context"
	"database/sql"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoRepoImpl) GetTodoList(ctx context.Context, id int) ([]entitytodo.TodoDetails, error) {
	functionName := "TodoRepoImpl.GetTodoList"

	var newQuery string
	var rows *sql.Rows
	var err error

	if id == 0 {
		newQuery = QueryGetTodoList + " WHERE deleted_at IS NULL;"
		rows, err = t.db.QueryContext(ctx, newQuery)
		if err != nil {
			t.l.Errorf("[%s] - In SELECT activity list query %s", functionName, err)
			return nil, cErrors.GetError(cErrors.InternalServer, err)
		}
		defer rows.Close()
	} else {
		newQuery = QueryGetTodoList + " WHERE activity_group_id=? AND IS NULL;"
		rows, err = t.db.QueryContext(ctx, newQuery, id)
		if err != nil {
			t.l.Errorf("[%s] - In SELECT activity list query %s", functionName, err)
			return nil, cErrors.GetError(cErrors.InternalServer, err)
		}
		defer rows.Close()
	}

	todos := make([]entitytodo.TodoDetails, 0)

	for rows.Next() {
		var todo entitytodo.TodoDetails
		if err = rows.Scan(
			&todo.Id,
			&todo.ActivityGroupId,
			&todo.Title,
			&todo.IsActive,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		); err != nil {
			t.l.Errorf("[%s] While scanning activity detail rows in todo list : %s", functionName, err)
			return nil, cErrors.GetError(cErrors.InternalServer, err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
