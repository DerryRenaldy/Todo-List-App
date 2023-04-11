package activityservices

import (
	"context"
	activityrepo "github.com/DerryRenaldy/Todo-List-App/apis/v1/repositories/activity"
	dtoactivity "github.com/DerryRenaldy/Todo-List-App/dto/activity"
	entityactivity "github.com/DerryRenaldy/Todo-List-App/entities/activity"
	"github.com/DerryRenaldy/logger/logger"
)

type ActivityServiceImpl struct {
	activityRepo activityrepo.IRepository
	l            logger.ILogger
}

func NewActivityServiceImpl(activityRepo activityrepo.IRepository, l logger.ILogger) *ActivityServiceImpl {
	return &ActivityServiceImpl{
		activityRepo: activityRepo,
		l:            l,
	}
}

type IService interface {
	CreateActivity(ctx context.Context, payload *dtoactivity.CreateActivityRequest) (*entityactivity.ActivityDetails, error)
	GetOneActivityById(ctx context.Context, id int) (*entityactivity.ActivityDetails, error)
	GetListActivity(ctx context.Context) ([]entityactivity.ActivityDetails, error)
	DeleteActivityById(ctx context.Context, id int) error
}

var _ IService = (*ActivityServiceImpl)(nil)
