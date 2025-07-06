package shade

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/shade/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *model.Shade) (*model.Shade, error)
	GetAll() ([]*model.Shade, error)
	GetById(id uint) (*model.Shade, error)
	Update(model *model.Shade) (*model.Shade, error)
	Delete(id uint) error
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *model.Shade) (*model.Shade, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll() ([]*model.Shade, error) {
	var shades []*model.Shade
	var err error

	result := r.DB.Find(&shades)
	if result.Error != nil {
		err = result.Error
	}

	return shades, err
}

func (r *repository) GetById(id uint) (*model.Shade, error) {
	var shade *model.Shade
	var err error

	result := r.DB.First(&shade, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return shade, err
}

func (r *repository) Update(model *model.Shade) (*model.Shade, error) {
	result := r.DB.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Delete(id uint) error {
	result := r.DB.Delete(&model.Shade{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
