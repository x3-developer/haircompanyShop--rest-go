package docsResponse

import (
	"haircompany-shop-rest/internal/modules/v1/category/dto"
)

type dashboardUserErrorField struct {
	Field     string `json:"field" enums:"name,slug,parentId"`
	ErrorCode string `json:"errorCode" enums:"NOT_UNIQUE,NOT_FOUND"`
}

type DashboardUserCreate201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type DashboardUserCreate400 struct {
	Response400
	Fields []dashboardUserErrorField `json:"fields,omitempty"`
}
