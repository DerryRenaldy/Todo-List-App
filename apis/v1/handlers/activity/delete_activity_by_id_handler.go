package activityhandlers

import (
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *ActivityHandlerImpl) DeleteActivityById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	deleteActivityErr := a.activityService.DeleteActivityById(ctx, idInt)
	if deleteActivityErr != nil {
		return deleteActivityErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    struct{}{},
	}

	w.WriteHeader(http.StatusOK)
	return handlers.ResponseJson(w, res)
}
