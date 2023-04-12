package todohandlers

import (
	todoservices "github.com/DerryRenaldy/Todo-List-App/apis/v1/services/todo"
	"github.com/DerryRenaldy/logger/logger"
	"net/http"
)

type TodoHandlerImpl struct {
	todoService todoservices.IService
	l           logger.ILogger
}

func NewTodoHandlerImpl(todoServices todoservices.IService, l logger.ILogger) *TodoHandlerImpl {
	return &TodoHandlerImpl{
		todoService: todoServices,
		l:           l,
	}
}

type IHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request) error
	GetOneTodoById(w http.ResponseWriter, r *http.Request) error
	GetTodoList(w http.ResponseWriter, r *http.Request) error
	DeleteTodoById(w http.ResponseWriter, r *http.Request) error
	UpdateTodoById(w http.ResponseWriter, r *http.Request) error
}

var _ IHandler = (*TodoHandlerImpl)(nil)
