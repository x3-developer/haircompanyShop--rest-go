package dto

import "haircompany-shop-rest/internal/modules/v1/product_type/model"

func TransformCreateDTOToModel(dto CreateDTO) *model.ProductType {
	return &model.ProductType{
		Name: dto.Name,
	}
}

func TransformUpdateDTOToModel(dto UpdateDTO, model *model.ProductType) *model.ProductType {
	if dto.Name != nil && *dto.Name != "" {
		model.Name = *dto.Name
	}
	return model
}

func TransformModelToResponseDTO(model *model.ProductType) *ResponseDTO {
	return &ResponseDTO{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Name:      model.Name,
	}
}
