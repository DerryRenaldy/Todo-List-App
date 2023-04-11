package activityrepo

import (
	"context"
	dtoactivity "github.com/DerryRenaldy/Todo-List-App/dto/activity"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	activitystore "github.com/DerryRenaldy/Todo-List-App/stores/mysql/activity"
)

type IRepository interface {
	CreateActivity(ctx context.Context, payload *dtoactivity.CreateActivityRequest) (*entityactivity.CreateActivityDetails, error)
}

var _ IRepository = (*activitystore.ActivityRepoImpl)(nil)
