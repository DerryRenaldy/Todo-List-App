package todoservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"time"
)

func (t *TodoServiceImpl) DeleteTodoById(ctx context.Context, id int) error {
	now := time.Now()

	_, getActivityErr := t.activityRepo.GetOneActivityById(ctx, id)
	if getActivityErr != nil {
		if getActivityErr.Error() == sql.ErrNoRows.Error() {
			return cErrors.GetError(cErrors.NotFound, errors.New(fmt.Sprintf("Activity with ID %d Not Found", id)))
		}
		return cErrors.GetError(cErrors.InternalServer, getActivityErr)
	}

	errDeleteTodo := t.todoRepo.DeleteTodoById(ctx, id, now)
	if errDeleteTodo != nil {
		return errDeleteTodo
	}

	return nil
}
