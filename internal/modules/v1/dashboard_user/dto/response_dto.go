package dto

import (
	"time"
)

type ResponseDTO struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}
