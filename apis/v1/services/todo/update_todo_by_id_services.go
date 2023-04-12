package todoservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoServiceImpl) UpdateTodoById(ctx context.Context, payload *dtotodo.UpdateTodoRequest) (*entitytodo.TodoDetails, error) {
	_, getActivityErr := t.todoRepo.GetOneTodoById(ctx, payload.Id)
	if getActivityErr != nil {
		if getActivityErr.Error() == sql.ErrNoRows.Error() {
			return nil, cErrors.GetError(cErrors.NotFound, errors.New(fmt.Sprintf("todo with ID %d Not Found", payload.Id)))
		}
		return nil, cErrors.GetError(cErrors.InternalServer, getActivityErr)
	}

	todoDetail, updateTodoErr := t.todoRepo.UpdateTodoById(ctx, payload)
	if updateTodoErr != nil {
		return nil, updateTodoErr
	}

	return todoDetail, nil
}
