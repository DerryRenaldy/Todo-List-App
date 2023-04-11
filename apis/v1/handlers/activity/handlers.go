package activityhandlers

import (
	activityservices "github.com/DerryRenaldy/Todo-List-App/apis/v1/services/activity"
	"github.com/DerryRenaldy/logger/logger"
	"net/http"
)

type ActivityHandlerImpl struct {
	activityService activityservices.IService
	l               logger.ILogger
}

func NewActivityHandlerImpl(activityService activityservices.IService, l logger.ILogger) *ActivityHandlerImpl {
	return &ActivityHandlerImpl{
		activityService: activityService,
		l:               l,
	}
}

type IHandler interface {
	CreateActivity(w http.ResponseWriter, r *http.Request) error
}

var _ IHandler = (*ActivityHandlerImpl)(nil)
