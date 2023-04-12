package todostore

import (
	"context"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"time"
)

func (t *TodoRepoImpl) DeleteTodoById(ctx context.Context, id int, deletedAt time.Time) error {
	functionName := "TodoRepoImpl.DeleteTodoById"

	_, err := t.db.ExecContext(ctx, QueryUpdateDeletedTime, deletedAt, id)
	if err != nil {
		t.l.Debugf("[%s] - While executing UPDATE statement to UPDATE todo : %s", functionName, err)
		return cErrors.GetError(cErrors.InternalServer, err)
	}

	return nil
}
