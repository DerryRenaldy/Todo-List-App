package todostore

import (
	"context"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoRepoImpl) UpdateTodoById(ctx context.Context, payload *dtotodo.UpdateTodoRequest) (*entitytodo.TodoDetails, error) {
	functionName := "TodoRepoImpl.UpdateTodoById"

	_, err := t.db.ExecContext(ctx, QueryUpdateTodoById, payload.Title, payload.Priority, payload.IsActive, payload.Id)
	if err != nil {
		t.l.Debugf("[%s] - While executing UPDATE statement to UPDATE activity : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	todoDetail, err := t.GetOneTodoById(ctx, payload.Id)
	if err != nil {
		t.l.Debugf("[%s] - While executing GET statement to GET activity by id : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	return todoDetail, nil
}
