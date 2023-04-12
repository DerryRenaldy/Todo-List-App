package todoservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (t *TodoServiceImpl) GetOneTodoById(ctx context.Context, id int) (*entitytodo.TodoDetails, error) {
	todoDetails, getOneTodoErr := t.todoRepo.GetOneTodoById(ctx, id)
	if getOneTodoErr != nil {
		if getOneTodoErr.Error() == sql.ErrNoRows.Error() {
			return nil, cErrors.GetError(cErrors.NotFound, errors.New(fmt.Sprintf("Todo with ID %d Not Found", id)))
		}
		return nil, getOneTodoErr
	}

	return todoDetails, nil
}
