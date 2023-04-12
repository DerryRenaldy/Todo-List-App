package todohandlers

import (
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"net/http"
	"regexp"
	"strconv"
)

func (t *TodoHandlerImpl) GetTodoList(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	id := r.URL.Query().Get("activity_group_id")
	var idInt int

	pattern := "^[0-9]+$"
	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(id) {
		idInt = 0
	}

	idInt, _ = strconv.Atoi(id)

	todoDetails, getTodoListErr := t.todoService.GetTodoList(ctx, idInt)
	if getTodoListErr != nil {
		return getTodoListErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoDetails,
	}

	w.WriteHeader(http.StatusOK)
	return handlers.ResponseJson(w, res)
}
