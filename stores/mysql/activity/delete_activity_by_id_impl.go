package activitystore

import (
	"context"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"time"
)

func (a *ActivityRepoImpl) DeleteActivityById(ctx context.Context, id int, deletedAt time.Time) error {
	functionName := "ActivityRepoImpl.DeleteActivityById"

	_, err := a.db.ExecContext(ctx, QueryUpdateDeletedTime, deletedAt, id)
	if err != nil {
		a.l.Debugf("[%s] - While executing UPDATE statement to UPDATE activity : %s", functionName, err)
		return cErrors.GetError(cErrors.InternalServer, err)
	}

	return nil
}
