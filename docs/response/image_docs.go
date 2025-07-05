package docsResponse

import "haircompany-shop-rest/internal/modules/v1/image/dto"

type imageErrorField struct {
	Field     string `json:"field" example:"images"`
	ErrorCode string `json:"errorCode" example:"INVALID_FORMAT"`
}

type ImageUpload200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type ImageUpload400 struct {
	Response400
	Fields []imageErrorField `json:"fields,omitempty"`
}
