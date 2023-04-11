package server

import (
	"database/sql"
	activityhandlers "github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers/activity"
	activityservices "github.com/DerryRenaldy/Todo-List-App/apis/v1/services/activity"
	"github.com/DerryRenaldy/Todo-List-App/configs"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/pkgs/database/mysql"
	"github.com/DerryRenaldy/Todo-List-App/server/middleware"
	activitystore "github.com/DerryRenaldy/Todo-List-App/stores/mysql/activity"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Server struct {
	cfg     *configs.Config
	log     logger.ILogger
	service activityservices.IService
	handler activityhandlers.IHandler
}

var addr string
var SVR *Server
var db *sql.DB
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func (s *Server) Register() {
	// Initiate new MySQL connection
	dbConn := mysql.NewConnection(s.log)
	if dbConn == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}
	db = dbConn.DBConnect()
	if db == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	activityRepo := activitystore.NewActivityRepoImpl(db, s.log)

	// Register service
	s.service = activityservices.NewActivityServiceImpl(activityRepo, s.log)

	// Register handler
	s.handler = activityhandlers.NewActivityHandlerImpl(s.service, s.log)
}

func NewService(cfg *configs.Config, logger logger.ILogger) *Server {
	if SVR != nil {
		return SVR
	}

	SVR = &Server{
		cfg: cfg,
		log: logger,
	}

	SVR.Register()

	return SVR
}

func (s *Server) Start() {
	addr = ":8090"

	r := mux.NewRouter()
	r.Handle(constants.GetListActivityEndpoint, middleware.ErrHandler(s.handler.GetListActivity)).Methods(http.MethodGet)
	r.Handle(constants.CreateActivityEndpoint, middleware.ErrHandler(s.handler.CreateActivity)).Methods(http.MethodPost)
	r.Handle(constants.GetOneActivityByIdEndpoint, middleware.ErrHandler(s.handler.GetOneActivityById)).Methods(http.MethodGet)
	r.Handle(constants.DeleteActivityByIdEndpoint, middleware.ErrHandler(s.handler.DeleteActivityById)).Methods(http.MethodDelete)
	r.Handle(constants.UpdateTitleActivityByIdEndpoint, middleware.ErrHandler(s.handler.UpdateTitleActivityById)).Methods(http.MethodPatch)

	s.log.Infof("HTTP server starting %v", addr)

	go func() {
		err := http.ListenAndServe(addr, r)
		if err != nil {
			s.log.Fatalf("error listening to address %v, err=%v", addr, err)
		}
		s.log.Infof("HTTP server started %v", addr)
	}()

	sig := <-signalChan
	s.log.Infof("%s signal caught", sig)

	// Doing cleanup if received signal from Operating System.
	err := db.Close()
	if err != nil {
		s.log.Errorf("Error in closing DB connection. Err : %+v", err.Error())
	}
}
