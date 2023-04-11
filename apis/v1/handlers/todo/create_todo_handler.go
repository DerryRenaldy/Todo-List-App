package todohandlers

import (
	"encoding/json"
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"io"
	"net/http"
)

func (t *TodoHandlerImpl) CreateTodo(w http.ResponseWriter, r *http.Request) error {
	functionName := "TodoHandlerImpl.CreateTodo"
	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	payload := dtotodo.CreateTodoRequest{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.l.Errorf("[%s : io.ReadAll - error reading request body] : %s", functionName, err)
		return cErrors.GetError(cErrors.InternalServer, err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		t.l.Errorf("[%s : json.Unmarshal - error parsing request body to struct] : %s", functionName, err)
		return cErrors.GetError(cErrors.InternalServer, err)
	}

	todoDetails, createTodoErr := t.todoService.CreateTodo(ctx, &payload)
	if createTodoErr != nil {
		return createTodoErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todoDetails,
	}

	w.WriteHeader(http.StatusCreated)
	return handlers.ResponseJson(w, res)
}

