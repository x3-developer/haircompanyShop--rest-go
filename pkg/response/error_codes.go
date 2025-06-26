package response

type ErrorCode string

const (
	BadRequest        ErrorCode = "BAD_REQUEST"
	ServerError       ErrorCode = "SERVER_ERROR"
	NotUnique         ErrorCode = "NOT_UNIQUE"
	MethodNotAllowed  ErrorCode = "METHOD_NOT_ALLOWED"
	NotFound          ErrorCode = "NOT_FOUND"
	RequestTooLarge   ErrorCode = "REQUEST_TOO_LARGE"
	FileTooLarge      ErrorCode = "FILE_TOO_LARGE"
	InvalidFileType   ErrorCode = "INVALID_FILE_TYPE"
	NotBlank          ErrorCode = "NOT_BLANK"
	MinLength         ErrorCode = "MIN_LENGTH"
	MaxLength         ErrorCode = "MAX_LENGTH"
	HasLinkedEntities ErrorCode = "HAS_LINKED_ENTITIES"
	Forbidden         ErrorCode = "FORBIDDEN"
)

func GetErrorCodeByTag(tag string) ErrorCode {
	switch tag {
	case "required":
		return NotBlank
	case "min":
		return MinLength
	case "max":
		return MaxLength
	default:
		return ServerError
	}
}
