package activitystore

import (
	"database/sql"
	"github.com/DerryRenaldy/logger/logger"
)

type ActivityRepoImpl struct {
	db *sql.DB
	l  logger.ILogger
}

func NewActivityRepoImpl(db *sql.DB, l logger.ILogger) *ActivityRepoImpl {
	return &ActivityRepoImpl{
		db: db,
		l:  l,
	}
}
