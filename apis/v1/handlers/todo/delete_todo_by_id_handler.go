package todohandlers

import (
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (t *TodoHandlerImpl) DeleteTodoById(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	deleteTodoErr := t.todoService.DeleteTodoById(ctx, idInt)
	if deleteTodoErr != nil {
		return deleteTodoErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    struct{}{},
	}

	w.WriteHeader(http.StatusOK)
	return handlers.ResponseJson(w, res)
}
