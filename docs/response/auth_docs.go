package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/auth/dto"
)

type authErrorField struct {
	Field     string `json:"field" enums:"email,password,refreshToken"`
	ErrorCode string `json:"errorCode" enums:"REQUIRED,INVALID_EMAIL,MIN_LENGTH,MAX_LENGTH,INVALID_TOKEN"`
}

type DashboardLogin200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type DashboardLogin400 struct {
	Response400
	Fields []authErrorField `json:"fields,omitempty"`
}

type DashboardRefreshToken200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type DashboardRefreshToken400 struct {
	Response400
	Fields []authErrorField `json:"fields,omitempty"`
}
