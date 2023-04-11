package activityservices

import (
	"context"
	dtoactivity "github.com/DerryRenaldy/Todo-List-App/dto/activity"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
)

func (a *ActivityServiceImpl) CreateActivity(ctx context.Context, payload *dtoactivity.CreateActivityRequest) (*entityactivity.CreateActivityDetails, error) {
	activityDetail, errInsertActivity := a.activityRepo.CreateActivity(ctx, payload)
	if errInsertActivity != nil {
		return nil, errInsertActivity
	}

	return activityDetail, nil
}
