package activityhandlers

import (
	"encoding/json"
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	dtoactivity "github.com/DerryRenaldy/Todo-List-App/dto/activity"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"io"
	"net/http"
)

func (a *ActivityHandlerImpl) CreateActivity(w http.ResponseWriter, r *http.Request) error {
	functionName := "ActivityHandlerImpl.CreateActivity"
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	payload := dtoactivity.CreateActivityRequest{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		a.l.Errorf("[%s : io.ReadAll - error reading request body] : %s", functionName, err)
		return cErrors.GetError(cErrors.InternalServer, err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		a.l.Errorf("[%s : json.Unmarshal - error parsing request body to struct] : %s", functionName, err)
		return cErrors.GetError(cErrors.InternalServer, err)
	}

	activityDetail, createActivityErr := a.activityService.CreateActivity(ctx, &payload)
	if createActivityErr != nil {
		return createActivityErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activityDetail,
	}

	w.WriteHeader(http.StatusCreated)
	return handlers.ResponseJson(w, res)
}
