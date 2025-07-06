package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/shade/dto"
)

type shadeErrorField struct {
	Field     string `json:"field" enums:"name,slug,parentId"`
	ErrorCode string `json:"errorCode" enums:"NOT_UNIQUE,NOT_FOUND"`
}

type ShadeCreate201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ShadeCreate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type ShadeList200 struct {
	IsSuccess bool              `json:"isSuccess" example:"true"`
	Data      []dto.ResponseDTO `json:"data"`
}

type ShadeGetById200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ShadeUpdate200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ShadeUpdate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type ShadeDelete200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}
