package dto

import "haircompany-shop-rest/internal/modules/v1/shade/model"

func TransformCreateDTOToModel(dto CreateDTO) *model.Shade {
	return &model.Shade{
		Name:      dto.Name,
		Image:     dto.Image,
		SortIndex: dto.SortIndex,
	}
}

func TransformUpdateDTOToModel(dto UpdateDTO, model *model.Shade) *model.Shade {
	if dto.Name != nil && *dto.Name != "" {
		model.Name = *dto.Name
	}
	if dto.Image != nil && *dto.Image != "" {
		model.Image = *dto.Image
	}
	if dto.SortIndex != nil {
		model.SortIndex = *dto.SortIndex
	}

	return model
}

func TransformModelToResponseDTO(model *model.Shade) *ResponseDTO {
	return &ResponseDTO{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Name:      model.Name,
		Image:     model.Image,
		SortIndex: model.SortIndex,
	}
}
