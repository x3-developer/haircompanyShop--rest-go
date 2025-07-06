package dto

type UpdateDTO struct {
	Name      *string `json:"name" validate:"omitempty,min=3,max=255"`
	Image     *string `json:"image" validate:"omitempty"`
	SortIndex *int    `json:"sort_index" validate:"omitempty,min=0"`
}
