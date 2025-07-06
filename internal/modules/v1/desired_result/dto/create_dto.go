package dto

type CreateDTO struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}
