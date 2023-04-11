package todostore

import (
	"database/sql"
	"github.com/DerryRenaldy/logger/logger"
)

type TodoRepoImpl struct {
	db *sql.DB
	l  logger.ILogger
}

func NewTodoRepoImpl(db *sql.DB, l logger.ILogger) *TodoRepoImpl {
	return &TodoRepoImpl{
		db: db,
		l:  l,
	}
}
