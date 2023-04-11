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

func (t *TodoServiceImpl) CreateTodo(ctx context.Context, payload *dtotodo.CreateTodoRequest) (*entitytodo.TodoDetails, error) {
	payload.IsActive = true
	payload.Priority = "very-high"

	_, getActivityErr := t.activityRepo.GetOneActivityById(ctx, payload.ActivityGroupId)
	if getActivityErr != nil {
		if getActivityErr.Error() == sql.ErrNoRows.Error() {
			return nil, cErrors.GetError(cErrors.NotFound, errors.New(
				fmt.Sprintf("Activity with ID %d Not Found", payload.ActivityGroupId)))
		}
		return nil, cErrors.GetError(cErrors.InternalServer, getActivityErr)
	}

	todoDetails, insertTodoErr := t.todoRepo.CreateTodo(ctx, payload)
	if insertTodoErr != nil {
		return nil, insertTodoErr
	}

	return todoDetails, nil
}
