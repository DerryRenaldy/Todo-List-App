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
	CreateTodoEndpoint = "/todo-items"
)

// Header library
const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	MIMEApplicationJson = "application/json"
)
