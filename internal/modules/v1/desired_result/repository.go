package desired_result

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/desired_result/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *model.DesiredResult) (*model.DesiredResult, error)
	GetAll() ([]*model.DesiredResult, error)
	GetById(id uint) (*model.DesiredResult, error)
	Update(model *model.DesiredResult) (*model.DesiredResult, error)
	Delete(id uint) error
	GetByUniqueFields(name string) (*model.DesiredResult, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *model.DesiredResult) (*model.DesiredResult, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll() ([]*model.DesiredResult, error) {
	var desiredResults []*model.DesiredResult
	var err error

	result := r.DB.Find(&desiredResults)
	if result.Error != nil {
		err = result.Error
	}

	return desiredResults, err
}

func (r *repository) GetById(id uint) (*model.DesiredResult, error) {
	var desiredResult *model.DesiredResult
	var err error

	result := r.DB.First(&desiredResult, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return desiredResult, err
}

func (r *repository) Update(model *model.DesiredResult) (*model.DesiredResult, error) {
	result := r.DB.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Delete(id uint) error {
	result := r.DB.Delete(&model.DesiredResult{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetByUniqueFields(name string) (*model.DesiredResult, error) {
	var desiredResult *model.DesiredResult
	var err error

	result := r.DB.First(&desiredResult, "name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return desiredResult, err
}
