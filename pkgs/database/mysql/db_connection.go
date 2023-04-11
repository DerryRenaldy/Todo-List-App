package mysql

import (
	"database/sql"
	"fmt"
	"github.com/DerryRenaldy/Todo-List-App/configs"
	"github.com/DerryRenaldy/logger/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

// config is configuration struct for db connection
type config struct {
	Database configs.MySQL
}

type Connection struct {
	Cfg *config
	Log logger.ILogger
}

func NewConnection(l logger.ILogger) *Connection {
	dbConfig := &config{
		Database: configs.MySQL{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Username: os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			DBName:   os.Getenv("MYSQL_DBNAME"),
		},
	}
	return &Connection{
		Log: l,
		Cfg: dbConfig,
	}
}

func (db *Connection) DBConnect() *sql.DB {
	dnsAddress := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db.Cfg.Database.Username, db.Cfg.Database.Password,
		db.Cfg.Database.Host, db.Cfg.Database.Port, db.Cfg.Database.DBName)

	dbConn, errConn := sql.Open(
		"mysql", dnsAddress)
	if errConn != nil {
		db.Log.Errorf("[ERR] Error while connecting... := %v", errConn)
		return nil
	}
	for dbConn.Ping() != nil {
		db.Log.Info("Attempting connect to DB...")
		time.Sleep(5 * time.Second)
	}

	return dbConn
}
