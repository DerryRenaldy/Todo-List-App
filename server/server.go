package server

import (
	"database/sql"
	activityhandlers "github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers/activity"
	todohandlers "github.com/DerryRenaldy/Todo-List-App/apis/v1/handlers/todo"
	activityservices "github.com/DerryRenaldy/Todo-List-App/apis/v1/services/activity"
	todoservices "github.com/DerryRenaldy/Todo-List-App/apis/v1/services/todo"
	"github.com/DerryRenaldy/Todo-List-App/configs"
	"github.com/DerryRenaldy/Todo-List-App/constants"
	"github.com/DerryRenaldy/Todo-List-App/pkgs/database/mysql"
	"github.com/DerryRenaldy/Todo-List-App/server/middleware"
	activitystore "github.com/DerryRenaldy/Todo-List-App/stores/mysql/activity"
	todostore "github.com/DerryRenaldy/Todo-List-App/stores/mysql/todo"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Server struct {
	cfg             *configs.Config
	log             logger.ILogger
	serviceActivity activityservices.IService
	serviceTodo     todoservices.IService
	handlerActivity activityhandlers.IHandler
	handlerTodo     todohandlers.IHandler
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
	todoRepo := todostore.NewTodoRepoImpl(db, s.log)

	// Register service
	s.serviceActivity = activityservices.NewActivityServiceImpl(activityRepo, s.log)
	s.serviceTodo = todoservices.NewTodoServiceImpl(activityRepo, todoRepo, s.log)

	// Register handler
	s.handlerActivity = activityhandlers.NewActivityHandlerImpl(s.serviceActivity, s.log)
	s.handlerTodo = todohandlers.NewTodoHandlerImpl(s.serviceTodo, s.log)
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

	// Activity route
	r.Handle(constants.GetListActivityEndpoint, middleware.ErrHandler(s.handlerActivity.GetListActivity)).
		Methods(http.MethodGet)
	r.Handle(constants.CreateActivityEndpoint, middleware.ErrHandler(s.handlerActivity.CreateActivity)).
		Methods(http.MethodPost)
	r.Handle(constants.GetOneActivityByIdEndpoint, middleware.ErrHandler(s.handlerActivity.GetOneActivityById)).
		Methods(http.MethodGet)
	r.Handle(constants.DeleteActivityByIdEndpoint, middleware.ErrHandler(s.handlerActivity.DeleteActivityById)).
		Methods(http.MethodDelete)
	r.Handle(constants.UpdateTitleActivityByIdEndpoint, middleware.ErrHandler(s.handlerActivity.UpdateTitleActivityById)).
		Methods(http.MethodPatch)

	// Todo route
	r.Handle(constants.CreateTodoEndpoint, middleware.ErrHandler(s.handlerTodo.CreateTodo)).
		Methods(http.MethodPost)
	r.Handle(constants.GetOneTodoByIdEndpoint, middleware.ErrHandler(s.handlerTodo.GetOneTodoById)).
		Methods(http.MethodGet)
	r.Handle(constants.GetTodoListEndpoint, middleware.ErrHandler(s.handlerTodo.GetTodoList)).
		Methods(http.MethodGet)

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
