package activitystore

import (
	"context"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (a *ActivityRepoImpl) GetListActivity(ctx context.Context) ([]entityactivity.ActivityDetails, error) {
	functionName := "ActivityRepoImpl.GetListActivity"

	rows, err := a.db.QueryContext(ctx, QueryGetListActivity)
	if err != nil {
		a.l.Errorf("[%s] - In SELECT activity list query %s", functionName, err)
		return nil, cErrors.GetError(cErrors.InternalServer, err)
	}
	defer rows.Close()
	activities := make([]entityactivity.ActivityDetails, 0)

	for rows.Next() {
		var activity entityactivity.ActivityDetails
		if err = rows.Scan(
			&activity.Id,
			&activity.Title,
			&activity.Email,
			&activity.CreatedAt,
			&activity.UpdatedAt,
		); err != nil {
			a.l.Errorf("[%s] While scanning activity detail rows in activity list : %s", functionName, err)
			return nil, cErrors.GetError(cErrors.InternalServer, err)
		}
		activities = append(activities, activity)
	}

	return activities, nil
}
