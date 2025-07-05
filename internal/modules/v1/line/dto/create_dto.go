package dto

type CreateDTO struct {
	Name  string `json:"name" validate:"required,min=3,max=255"`
	Color string `json:"color" validate:"required,hex_color" example:"#FF5733"`
}
