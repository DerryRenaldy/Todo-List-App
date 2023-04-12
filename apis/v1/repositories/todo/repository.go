package todorepo

import (
	"context"
	dtotodo "github.com/DerryRenaldy/Todo-List-App/dto/todo"
	entitytodo "github.com/DerryRenaldy/Todo-List-App/entities/todo"
	todostore "github.com/DerryRenaldy/Todo-List-App/stores/mysql/todo"
)

type IRepository interface {
	CreateTodo(ctx context.Context, payload *dtotodo.CreateTodoRequest) (*entitytodo.TodoDetails, error)
	GetOneTodoById(ctx context.Context, id int) (*entitytodo.TodoDetails, error)
	GetTodoList(ctx context.Context, id int) ([]entitytodo.TodoDetails, error)
}

var _ IRepository = (*todostore.TodoRepoImpl)(nil)
