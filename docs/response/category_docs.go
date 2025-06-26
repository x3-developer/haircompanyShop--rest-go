package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/category/dto"
)

type createField struct {
	Field     string `json:"field" enums:"name,slug,parentId"`
	ErrorCode string `json:"errorCode" enums:"NOT_UNIQUE,NOT_FOUND"`
}

type Category201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type Category500 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode" enums:"SERVER_ERROR"`
}

type CategoryCreate400 struct {
	IsSuccess bool          `json:"isSuccess" example:"false"`
	Message   string        `json:"message"`
	ErrorCode string        `json:"errorCode" enums:"BAD_REQUEST"`
	Fields    []createField `json:"fields,omitempty"`
}

type CategoryList200 struct {
	IsSuccess bool              `json:"isSuccess" example:"true"`
	Data      []dto.ResponseDTO `json:"data"`
}

type CategoryGetById200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type CategoryGetById400 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode" enums:"BAD_REQUEST"`
}

type CategoryGetById404 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode" enums:"NOT_FOUND"`
}

type CategoryUpdate200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type CategoryUpdate400 struct {
	IsSuccess bool          `json:"isSuccess" example:"false"`
	Message   string        `json:"message"`
	ErrorCode string        `json:"errorCode" enums:"BAD_REQUEST"`
	Fields    []createField `json:"fields,omitempty"`
}

type CategoryDelete200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type CategoryDelete400 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode" enums:"BAD_REQUEST"`
}

type CategoryDelete404 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode" enums:"NOT_FOUND"`
}

type CategoryDelete409 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode" enums:"HAS_LINKED_ENTITIES"`
}
