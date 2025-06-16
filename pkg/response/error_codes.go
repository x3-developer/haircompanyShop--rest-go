package response

type ErrorCode string

const (
	BadRequest       ErrorCode = "BAD_REQUEST"
	ServerError      ErrorCode = "SERVER_ERROR"
	NotUnique        ErrorCode = "NOT_UNIQUE"
	MethodNotAllowed ErrorCode = "METHOD_NOT_ALLOWED"
	NotFound         ErrorCode = "NOT_FOUND"
)
