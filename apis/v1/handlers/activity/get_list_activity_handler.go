package activityhandlers

import (
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"net/http"
)

func (a *ActivityHandlerImpl) GetListActivity(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	activityDetail, GetActivityErr := a.activityService.GetListActivity(ctx)
	if GetActivityErr != nil {
		return GetActivityErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityDetail,
	}

	w.WriteHeader(http.StatusOK)
	return handlers.ResponseJson(w, res)
}
