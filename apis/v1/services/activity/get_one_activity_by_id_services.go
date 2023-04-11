package activityservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
)

func (a *ActivityServiceImpl) GetOneActivityById(ctx context.Context, id int) (*entityactivity.ActivityDetails, error) {
	activityDetail, errGetActivity := a.activityRepo.GetOneActivityById(ctx, id)
	if errGetActivity != nil {
		if errGetActivity.Error() == sql.ErrNoRows.Error() {
			return nil, cErrors.GetError(cErrors.NotFound, errors.New(fmt.Sprintf("Activity with ID %d Not Found", id)))
		}
		return nil, errGetActivity
	}

	return activityDetail, nil
}
