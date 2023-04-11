package activityrepo

import (
	"context"
	dtoactivity "github.com/DerryRenaldy/Todo-List-App/dto/activity"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	activitystore "github.com/DerryRenaldy/Todo-List-App/stores/mysql/activity"
	"time"
)

type IRepository interface {
	CreateActivity(ctx context.Context, payload *dtoactivity.CreateActivityRequest) (*entityactivity.ActivityDetails, error)
	GetOneActivityById(ctx context.Context, id int) (*entityactivity.ActivityDetails, error)
	GetListActivity(ctx context.Context) ([]entityactivity.ActivityDetails, error)
	DeleteActivityById(ctx context.Context, id int, deletedAt time.Time) error
	UpdateActivityTitleById(ctx context.Context, id int, title string) (*entityactivity.ActivityDetails, error)
}

var _ IRepository = (*activitystore.ActivityRepoImpl)(nil)
