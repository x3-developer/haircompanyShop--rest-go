package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/line/dto"
)

type lineErrorField struct {
	Field     string `json:"field" enums:"name,slug,parentId"`
	ErrorCode string `json:"errorCode" enums:"NOT_UNIQUE,NOT_FOUND"`
}

type LineCreate201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type LineCreate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type LineList200 struct {
	IsSuccess bool              `json:"isSuccess" example:"true"`
	Data      []dto.ResponseDTO `json:"data"`
}

type LineGetById200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type LineUpdate200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type LineUpdate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type LineDelete200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}
