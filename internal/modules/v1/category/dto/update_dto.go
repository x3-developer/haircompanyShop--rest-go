package dto

type UpdateDTO struct {
	Name            *string `json:"name" validate:"omitempty,min=3,max=255"`
	Description     *string `json:"description" validate:"omitempty"`
	Image           *string `json:"image" validate:"omitempty"`
	HeaderImage     *string `json:"headerImage" validate:"omitempty"`
	Slug            *string `json:"slug" validate:"omitempty,min=3,max=255"`
	ParentID        *uint   `json:"parentId" validate:"omitempty"`
	SortIndex       *int    `json:"sortIndex" validate:"omitempty,gte=0"`
	SeoTitle        *string `json:"seoTitle" validate:"omitempty"`
	SeoDescription  *string `json:"seoDescription" validate:"omitempty"`
	SeoKeys         *string `json:"seoKeys" validate:"omitempty"`
	IsActive        *bool   `json:"isActive" validate:"omitempty"`
	IsShade         *bool   `json:"isShade" validate:"omitempty"`
	IsVisibleInMenu *bool   `json:"isVisibleInMenu" validate:"omitempty"`
	IsVisibleOnMain *bool   `json:"isVisibleOnMain" validate:"omitempty"`
}
