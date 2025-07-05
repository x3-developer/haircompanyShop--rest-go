package dto

type UpdateDTO struct {
	Name  *string `json:"name" validate:"omitempty,min=3,max=255"`
	Color *string `json:"color" validate:"omitempty,hex_color"`
}
