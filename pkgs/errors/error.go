package cErrors

// List of all errors
const (
	InternalServer     = "internal error"
	InvalidRequest     = "Invalid request"
	UnavailableService = "Unavailable Service"
	NotFound           = "Not Found"
)

const (
	UnavailableServiceCode = 503
	InternalServerCode     = 500
	InvalidRequestCode     = 400
	NotFoundCode           = 404
)

type IError interface {
	Error() string
	GetHTTPCode() int
}

type CustomError struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c CustomError) Error() string {
	return c.Message
}

func (c CustomError) GetHTTPCode() int {
	return c.Code
}

// getErrorCode return http error code base on error message.
func getErrorCode(errMsg string) int {
	var val int
	switch errMsg {
	case InternalServer:
		val = InternalServerCode
	case UnavailableService:
		val = UnavailableServiceCode
	case InvalidRequest:
		val = InvalidRequestCode
	case NotFound:
		val = NotFoundCode
	default:
		val = InternalServerCode
	}
	return val
}

// GetError code and message then return.
func GetError(errMessage string, errDetail error) *CustomError {
	return &CustomError{
		Code:    getErrorCode(errMessage),
		Status:  errMessage,
		Message: errDetail.Error(),
	}
}
