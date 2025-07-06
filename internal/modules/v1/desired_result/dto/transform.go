package dto

import "haircompany-shop-rest/internal/modules/v1/desired_result/model"

func TransformCreateDTOToModel(dto CreateDTO) *model.DesiredResult {
	return &model.DesiredResult{
		Name: dto.Name,
	}
}

func TransformUpdateDTOToModel(dto UpdateDTO, model *model.DesiredResult) *model.DesiredResult {
	if dto.Name != nil && *dto.Name != "" {
		model.Name = *dto.Name
	}
	return model
}

func TransformModelToResponseDTO(model *model.DesiredResult) *ResponseDTO {
	return &ResponseDTO{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Name:      model.Name,
	}
}
