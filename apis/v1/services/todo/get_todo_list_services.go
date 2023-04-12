package todoservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoServiceImpl) GetTodoList(ctx context.Context, id int) ([]entitytodo.TodoDetails, error) {
	if id != 0 {
		_, getActivityErr := t.activityRepo.GetOneActivityById(ctx, id)
		if getActivityErr != nil {
			if getActivityErr.Error() == sql.ErrNoRows.Error() {
				return nil, cErrors.GetError(cErrors.NotFound, errors.New(
					fmt.Sprintf("Activity with ID %d Not Found", id)))
			}
			return nil, cErrors.GetError(cErrors.InternalServer, getActivityErr)
		}
	}

	todoDetails, getTodoListErr := t.todoRepo.GetTodoList(ctx, id)
	if getTodoListErr != nil {
		return nil, getTodoListErr
	}

	return todoDetails, nil
}
