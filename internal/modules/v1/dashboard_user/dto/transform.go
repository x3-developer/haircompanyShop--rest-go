package dto

import (
	dashboardUserModel "haircompany-shop-rest/internal/modules/v1/dashboard_user/model"
)

func TransformCreateDTOToModel(dto CreateDTO) *dashboardUserModel.DashboardUser {
	return &dashboardUserModel.DashboardUser{
		Email: dto.Email,
		Role:  dto.Role,
	}
}

func TransformModelToResponseDTO(model *dashboardUserModel.DashboardUser) *ResponseDTO {
	return &ResponseDTO{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Email:     model.Email,
		Role:      model.Role,
	}
}
