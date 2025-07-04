package dto

type CreateDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Role     string `json:"role" validate:"required,oneof=admin manager"`
}
