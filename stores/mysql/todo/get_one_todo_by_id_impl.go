package todostore

import (
	"context"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoRepoImpl) GetOneTodoById(ctx context.Context, id int) (*entitytodo.TodoDetails, error) {
	functionName := "ActivityRepoImpl.GetOneTodoById"

	todoDetails := entitytodo.TodoDetails{}
	err := t.db.QueryRowContext(ctx, QueryGetOneTodoById, id).
		Scan(
			&todoDetails.Id,
			&todoDetails.ActivityGroupId,
			&todoDetails.Title,
			&todoDetails.IsActive,
			&todoDetails.Priority,
			&todoDetails.CreatedAt,
			&todoDetails.UpdatedAt,
		)
	if err != nil {
		t.l.Debugf("[%s] - While executing GET statement to GET todo by id : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	return &todoDetails, nil
}
