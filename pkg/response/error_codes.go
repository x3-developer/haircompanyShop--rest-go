package response

type ErrorCode string

const (
	BadRequest       ErrorCode = "BAD_REQUEST"
	ServerError      ErrorCode = "SERVER_ERROR"
	NotUnique        ErrorCode = "NOT_UNIQUE"
	MethodNotAllowed ErrorCode = "METHOD_NOT_ALLOWED"
	NotFound         ErrorCode = "NOT_FOUND"
	RequestTooLarge  ErrorCode = "REQUEST_TOO_LARGE"
	FileTooLarge     ErrorCode = "FILE_TOO_LARGE"
	InvalidFileType  ErrorCode = "INVALID_FILE_TYPE"
)
