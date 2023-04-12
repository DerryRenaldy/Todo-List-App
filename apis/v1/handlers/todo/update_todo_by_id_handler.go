package todohandlers

import (
	"encoding/json"
	"github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	cErrors "github.com/DerryRenaldy/Todo-List-App/pkgs/errors"
	"github.com/DerryRenaldy/Todo-List-App/response"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func (t *TodoHandlerImpl) UpdateTodoById(w http.ResponseWriter, r *http.Request) error {
	functionName := "ActivityHandlerImpl.UpdateTitleActivityById"

	ctx := r.Context()
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJson)

	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)

	payload := dtotodo.UpdateTodoRequest{}
	payload.Id = idInt

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

	updateTodoDetail, updateTodoErr := t.todoService.UpdateTodoById(ctx, &payload)
	if updateTodoErr != nil {
		return updateTodoErr
	}

	res := response.CommonResponse{
		Status:  "Success",
		Message: "Success",
		Data:    updateTodoDetail,
	}

	w.WriteHeader(http.StatusOK)
	return handlers.ResponseJson(w, res)
}

