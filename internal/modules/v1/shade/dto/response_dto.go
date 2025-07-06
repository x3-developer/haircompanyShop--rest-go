package dto

import "time"

type ResponseDTO struct {
	Id        uint      `json:"id" example:"1"`
	CreatedAt time.Time `json:"createdAt" example:"2023-10-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-10-01T12:00:00Z"`
	Name      string    `json:"name" example:"Shade Name"`
	Image     string    `json:"image" example:"shade.png"`
	SortIndex int       `json:"sortIndex" example:"1"`
}
