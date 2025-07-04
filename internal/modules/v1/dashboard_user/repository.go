package dashboard_user

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *model.DashboardUser) (*model.DashboardUser, error)
	GetByEmail(email string) (*model.DashboardUser, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *model.DashboardUser) (*model.DashboardUser, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetByEmail(email string) (*model.DashboardUser, error) {
	var user *model.DashboardUser
	var err error

	result := r.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return user, err
}
