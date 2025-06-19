package dto

type UpdateDTO struct {
	Name            string `json:"name" validate:"min=3,max=255"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	HeaderImage     string `json:"headerImage"`
	Slug            string `json:"slug" validate:"min=3,max=255"`
	ParentID        *uint  `json:"parentId"`
	SortIndex       int    `json:"sortIndex" validate:"gte=0"`
	SeoTitle        string `json:"seoTitle"`
	SeoDescription  string `json:"seoDescription"`
	SeoKeys         string `json:"seoKeys"`
	IsActive        bool   `json:"isActive"`
	IsShade         bool   `json:"isShade"`
	IsVisibleInMenu bool   `json:"isVisibleInMenu"`
	IsVisibleOnMain bool   `json:"isVisibleOnMain"`
}
