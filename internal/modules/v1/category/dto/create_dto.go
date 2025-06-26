package dto

type CreateDTO struct {
	Name            string `json:"name" validate:"required,min=3,max=255"`
	Description     string `json:"description"`
	Image           string `json:"image" validate:"required,max=255"`
	HeaderImage     string `json:"headerImage" validate:"required,max=255"`
	Slug            string `json:"slug" validate:"required,min=3,max=255"`
	ParentID        *uint  `json:"parentId" validate:"omitempty,gte=0"` // Nullable, can be zero
	SortIndex       int    `json:"sortIndex" validate:"required,gte=0,lte=9999"`
	SeoTitle        string `json:"seoTitle" validate:"max=255"`
	SeoDescription  string `json:"seoDescription"`
	SeoKeys         string `json:"seoKeys"`
	IsActive        bool   `json:"isActive"`
	IsShade         bool   `json:"isShade"`
	IsVisibleInMenu bool   `json:"isVisibleInMenu"`
	IsVisibleOnMain bool   `json:"isVisibleOnMain"`
}
