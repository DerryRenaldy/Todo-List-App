package constants

// Endpoint Activity library
const (
	CreateActivityEndpoint          = "/activity-groups"
	GetOneActivityByIdEndpoint      = "/activity-groups/{id}"
	GetListActivityEndpoint         = "/activity-groups"
	DeleteActivityByIdEndpoint      = "/activity-groups/{id}"
	UpdateTitleActivityByIdEndpoint = "/activity-groups/{id}"
)

// Endpoint Todo library
const (
	CreateTodoEndpoint     = "/todo-items"
	GetOneTodoByIdEndpoint = "/todo-items/{id}"
	GetTodoListEndpoint    = "/todo-items"
	DeleteTodoByIdEndpoint = "/todo-items/{id}"
	UpdateTodoByIdEndpoint = "/todo-items/{id}"
)

// Header library
const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	MIMEApplicationJson = "application/json"
)
