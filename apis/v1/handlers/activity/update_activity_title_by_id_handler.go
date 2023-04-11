package activityhandlers

import (
	"encoding/json"
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func (a *ActivityHandlerImpl) UpdateTitleActivityById(w http.ResponseWriter, r *http.Request) error {
	functionName := "ActivityHandlerImpl.UpdateTitleActivityById"

	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	payload := struct {
		Title string `json:"title"`
	}{}

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

	updateActivityDetails, updateActivityErr := a.activityService.UpdateTitleActivityById(ctx, idInt, payload.Title)
	if updateActivityErr != nil {
		return updateActivityErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    updateActivityDetails,
	}

	w.WriteHeader(http.StatusOK)
	return handlers.ResponseJson(w, res)
}
