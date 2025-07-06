package dto

type CreateDTO struct {
	Name      string `json:"name" validate:"required,min=3,max=255"`
	Image     string `json:"image" validate:"required,max=255"`
	SortIndex int    `json:"sortIndex" validate:"required,gte=0,lte=9999"`
}
