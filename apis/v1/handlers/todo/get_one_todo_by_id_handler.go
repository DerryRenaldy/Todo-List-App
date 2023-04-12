package todohandlers

import (
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (t *TodoHandlerImpl) GetOneTodoById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	todoDetails, getOneTodoErr := t.todoService.GetOneTodoById(ctx, idInt)
	if getOneTodoErr != nil {
		return getOneTodoErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoDetails,
	}

	w.WriteHeader(http.StatusCreated)
	return handlers.ResponseJson(w, res)
}
