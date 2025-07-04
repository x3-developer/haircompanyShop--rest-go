package client_user

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/client_user/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	GetByPhone(phone string) (*model.ClientUser, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetByPhone(phone string) (*model.ClientUser, error) {
	var user *model.ClientUser
	var err error

	result := r.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return user, err
}
