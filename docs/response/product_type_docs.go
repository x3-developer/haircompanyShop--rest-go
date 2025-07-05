package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/product_type/dto"
)

type productTypeErrorField struct {
	Field     string `json:"field" enums:"name,slug,parentId"`
	ErrorCode string `json:"errorCode" enums:"NOT_UNIQUE,NOT_FOUND"`
}

type ProductTypeCreate201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ProductTypeCreate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type ProductTypeList200 struct {
	IsSuccess bool              `json:"isSuccess" example:"true"`
	Data      []dto.ResponseDTO `json:"data"`
}

type ProductTypeGetById200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ProductTypeUpdate200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ProductTypeUpdate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type ProductTypeDelete200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}
