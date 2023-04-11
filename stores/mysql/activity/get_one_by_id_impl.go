package activitystore

import (
	"context"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (a *ActivityRepoImpl) GetOneActivityById(ctx context.Context, id int) (*entityactivity.ActivityDetails, error) {
	functionName := "ActivityRepoImpl.GetOneById"

	activityDetail := entityactivity.ActivityDetails{}
	err := a.db.QueryRowContext(ctx, QueryGetSingleActivityById, id).
		Scan(
			&activityDetail.Id,
			&activityDetail.Title,
			&activityDetail.Email,
			&activityDetail.CreatedAt,
			&activityDetail.UpdatedAt,
		)
	if err != nil {
		a.l.Debugf("[%s] - While executing GET statement to GET activity by id : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	return &activityDetail, nil
}
