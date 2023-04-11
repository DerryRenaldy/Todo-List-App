package activityservices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"time"
)

func (a *ActivityServiceImpl) DeleteActivityById(ctx context.Context, id int) error {
	now := time.Now()

	_, getActivityErr := a.activityRepo.GetOneActivityById(ctx, id)
	if getActivityErr != nil {
		if getActivityErr.Error() == sql.ErrNoRows.Error() {
			return cErrors.GetError(cErrors.NotFound, errors.New(fmt.Sprintf("Activity with ID %d Not Found", id)))
		}
		return cErrors.GetError(cErrors.InternalServer, getActivityErr)
	}

	errDeleteActivity := a.activityRepo.DeleteActivityById(ctx, id, now)
	if errDeleteActivity != nil {
		return errDeleteActivity
	}

	return nil
}
