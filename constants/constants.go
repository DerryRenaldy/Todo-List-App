package constants

// Endpoint library
const (
	CreateActivityEndpoint          = "/activity-groups"
	GetOneActivityByIdEndpoint      = "/activity-groups/{id}"
	GetListActivityEndpoint         = "/activity-groups"
	DeleteActivityByIdEndpoint      = "/activity-groups/{id}"
	UpdateTitleActivityByIdEndpoint = "/activity-groups/{id}"
)

// Header library
const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	MIMEApplicationJson = "application/json"
)
