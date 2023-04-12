package todoservices

import (
	"context"
	activityrepo "github.com/DerryRenaldy/Todo-List-App/apis/v1/repositories/activity"
	todorepo "github.com/DerryRenaldy/Todo-List-App/apis/v1/repositories/todo"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	"github.com/DerryRenaldy/logger/logger"
)

type TodoServiceImpl struct {
	activityRepo activityrepo.IRepository
	todoRepo     todorepo.IRepository
	l            logger.ILogger
}

func NewTodoServiceImpl(activityRepo activityrepo.IRepository, todoRepo todorepo.IRepository,
	l logger.ILogger) *TodoServiceImpl {
	return &TodoServiceImpl{
		activityRepo: activityRepo,
		todoRepo:     todoRepo,
		l:            l,
	}
}

type IService interface {
	CreateTodo(ctx context.Context, payload *dtotodo.CreateTodoRequest) (*entitytodo.TodoDetails, error)
	GetOneTodoById(ctx context.Context, id int) (*entitytodo.TodoDetails, error)
	GetTodoList(ctx context.Context, id int) ([]entitytodo.TodoDetails, error)
	DeleteTodoById(ctx context.Context, id int) error
	UpdateTodoById(ctx context.Context, payload *dtotodo.UpdateTodoRequest) (*entitytodo.TodoDetails, error)
}

var _ IService = (*TodoServiceImpl)(nil)
