package docsResponse

type Response500 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Internal server error"`
	ErrorCode string `json:"errorCode" enums:"SERVER_ERROR"`
}

type Response400 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Bad request or validation error"`
	ErrorCode string `json:"errorCode" enums:"BAD_REQUEST"`
}

type Response401 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Unauthorized"`
	ErrorCode string `json:"errorCode" enums:"UNAUTHORIZED"`
}

type Response404 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Resource not found"`
	ErrorCode string `json:"errorCode" enums:"NOT_FOUND"`
}

type Response409 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Conflict"`
	ErrorCode string `json:"errorCode" enums:"HAS_LINKED_ENTITIES"`
}

type Response413 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Request entity too large"`
	ErrorCode string `json:"errorCode" enums:"REQUEST_TOO_LARGE"`
}
