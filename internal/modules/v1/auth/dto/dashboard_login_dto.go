package dto

type DashboardLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}
