package dto

import (
	categoryModel "haircompany-shop-rest/internal/modules/v1/category/model"
)

func TransformCreateDTOToModel(dto CreateDTO) *categoryModel.Category {
	return &categoryModel.Category{
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

func TransformUpdateDTOToModel(dto UpdateDTO, model *categoryModel.Category) *categoryModel.Category {
	if dto.Name != nil && *dto.Name != "" {
		model.Name = *dto.Name
	}
	if dto.Description != nil {
		model.Description = *dto.Description
	}
	if dto.Image != nil && *dto.Image != "" {
		model.Image = *dto.Image
	}
	if dto.HeaderImage != nil && *dto.HeaderImage != "" {
		model.HeaderImage = *dto.HeaderImage
	}
	if dto.Slug != nil && *dto.Slug != "" {
		model.Slug = *dto.Slug
	}
	if dto.ParentID != nil {
		model.ParentID = dto.ParentID
	}
	if dto.SortIndex != nil {
		model.SortIndex = *dto.SortIndex
	}
	if dto.SeoTitle != nil {
		model.SeoTitle = *dto.SeoTitle
	}
	if dto.SeoDescription != nil {
		model.SeoDescription = *dto.SeoDescription
	}
	if dto.SeoKeys != nil {
		model.SeoKeys = *dto.SeoKeys
	}
	if dto.IsActive != nil {
		model.IsActive = *dto.IsActive
	}
	if dto.IsShade != nil {
		model.IsShade = *dto.IsShade
	}
	if dto.IsVisibleInMenu != nil {
		model.IsVisibleInMenu = *dto.IsVisibleInMenu
	}
	if dto.IsVisibleOnMain != nil {
		model.IsVisibleOnMain = *dto.IsVisibleOnMain
	}

	return model
}

func TransformModelToResponseDTO(model *categoryModel.Category) *ResponseDTO {
	return &ResponseDTO{
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
