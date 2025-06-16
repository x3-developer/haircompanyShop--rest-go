package category

import "haircompany-shop-rest/internal/modules/v1/category/dto"

func TransformCreateDTOToModel(dto dto.CreateDTO) *Category {
	return &Category{
		Name:            dto.Name,
		Description:     dto.Description,
		Image:           dto.Image,
		HeaderImage:     dto.HeaderImage,
		Slug:            dto.Slug,
		ParentID:        dto.ParentID,
		SortIndex:       dto.SortIndex,
		SeoTitle:        dto.SeoTitle,
		SeoDescription:  dto.SeoDescription,
		SeoKeys:         dto.SeoKeys,
		IsActive:        dto.IsActive,
		IsShade:         dto.IsShade,
		IsVisibleInMenu: dto.IsVisibleInMenu,
		IsVisibleOnMain: dto.IsVisibleOnMain,
	}
}

func TransformModelToResponseDTO(model *Category) *dto.ResponseDTO {
	return &dto.ResponseDTO{
		Id:              model.ID,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		Name:            model.Name,
		Description:     model.Description,
		Image:           model.Image,
		HeaderImage:     model.HeaderImage,
		Slug:            model.Slug,
		ParentID:        model.ParentID,
		SortIndex:       model.SortIndex,
		SeoTitle:        model.SeoTitle,
		SeoDescription:  model.SeoDescription,
		SeoKeys:         model.SeoKeys,
		IsActive:        model.IsActive,
		IsShade:         model.IsShade,
		IsVisibleInMenu: model.IsVisibleInMenu,
		IsVisibleOnMain: model.IsVisibleOnMain,
	}
}
