package dto

import "haircompany-shop-rest/internal/modules/v1/line/model"

func TransformCreateDTOToModel(dto CreateDTO) *model.Line {
	return &model.Line{
		Name:  dto.Name,
		Color: dto.Color,
	}
}

func TransformUpdateDTOToModel(dto UpdateDTO, model *model.Line) *model.Line {
	if dto.Name != nil && *dto.Name != "" {
		model.Name = *dto.Name
	}
	if dto.Color != nil && *dto.Color != "" {
		model.Color = *dto.Color
	}
	return model
}

func TransformModelToResponseDTO(model *model.Line) *ResponseDTO {
	return &ResponseDTO{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Name:      model.Name,
		Color:     model.Color,
	}
}
