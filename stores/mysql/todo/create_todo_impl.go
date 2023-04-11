package todostore

import (
	"context"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoRepoImpl) CreateTodo(ctx context.Context, payload *dtotodo.CreateTodoRequest) (*entitytodo.TodoDetails, error) {
	functionName := "ActivityRepoImpl.CreateActivity"

	res, err := t.db.ExecContext(ctx, QueryCreateTodo, payload.ActivityGroupId, payload.Title, payload.Priority, payload.IsActive)
	if err != nil {
		t.l.Debugf("[%s] - While executing INSERT statement to INSERT todo : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	id, _ := res.LastInsertId()

	activityDetail, err := t.GetOneTodoById(ctx, int(id))
	if err != nil {
		t.l.Debugf("[%s] - While executing GET statement to GET todo by id : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	return activityDetail, nil
}
