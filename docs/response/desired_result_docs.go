package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/desired_result/dto"
)

type desiredResultErrorField struct {
	Field     string `json:"field" enums:"name,slug,parentId"`
	ErrorCode string `json:"errorCode" enums:"NOT_UNIQUE,NOT_FOUND"`
}

type DesiredResultCreate201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type DesiredResultCreate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type DesiredResultList200 struct {
	IsSuccess bool              `json:"isSuccess" example:"true"`
	Data      []dto.ResponseDTO `json:"data"`
}

type DesiredResultGetById200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type DesiredResultUpdate200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type DesiredResultUpdate400 struct {
	Response400
	Fields []categoryErrorField `json:"fields,omitempty"`
}

type DesiredResultDelete200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}
