package activitystore

import (
	"context"
	dtoactivity "github.com/DerryRenaldy/Todo-List-App/dto/activity"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	"github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (a *ActivityRepoImpl) CreateActivity(ctx context.Context, payload *dtoactivity.CreateActivityRequest) (*entityactivity.CreateActivityDetails, error) {
	functionName := "ActivityRepoImpl.CreateActivity"

	res, err := a.db.ExecContext(ctx, QueryCreateActivity, payload.Title, payload.Email)
	if err != nil {
		a.l.Debugf("[%s] - While executing INSERT statement to INSERT activity : %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}

	id, _ := res.LastInsertId()

	activityDetail := entityactivity.CreateActivityDetails{}
	err = a.db.QueryRowContext(ctx, QueryGetSingleActivityById, id).
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
