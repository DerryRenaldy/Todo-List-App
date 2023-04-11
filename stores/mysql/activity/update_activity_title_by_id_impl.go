package activitystore

import (
	"context"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (a *ActivityRepoImpl) UpdateActivityTitleById(ctx context.Context, id int, title string) (*entityactivity.ActivityDetails, error) {
	functionName := "ActivityRepoImpl.UpdateActivityTitleById"

	_, err := a.db.ExecContext(ctx, QueryUpdateTitleById, title, id)
	if err != nil {
		a.l.Debugf("[%s] - While executing UPDATE statement to UPDATE activity : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	activityDetail, err := a.GetOneActivityById(ctx, id)
	if err != nil {
		a.l.Debugf("[%s] - While executing GET statement to GET activity by id : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	return activityDetail, nil
}
