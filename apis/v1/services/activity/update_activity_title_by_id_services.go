package activityservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (a *ActivityServiceImpl) UpdateTitleActivityById(ctx context.Context, id int, title string) (*entityactivity.ActivityDetails, error) {
	_, getActivityErr := a.activityRepo.GetOneActivityById(ctx, id)
	if getActivityErr != nil {
		if getActivityErr.Error() == sql.ErrNoRows.Error() {
			return nil, cErrors.GetError(cErrors.NotFound, errors.New(fmt.Sprintf("Activity with ID %d Not Found", id)))
		}
		return nil, cErrors.GetError(cErrors.InternalServer, getActivityErr)
	}

	activityDetail, errDeleteActivity := a.activityRepo.UpdateActivityTitleById(ctx, id, title)
	if errDeleteActivity != nil {
		return nil, errDeleteActivity
	}

	return activityDetail, nil
}
