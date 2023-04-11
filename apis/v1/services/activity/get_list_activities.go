package activityservices

import (
	"context"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
)

func (a *ActivityServiceImpl) GetListActivity(ctx context.Context) ([]entityactivity.ActivityDetails, error) {
	activityDetail, errGetActivity := a.activityRepo.GetListActivity(ctx)
	if errGetActivity != nil {
		return nil, errGetActivity
	}

	return activityDetail, nil
}
