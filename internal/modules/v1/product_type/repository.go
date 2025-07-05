package product_type

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/product_type/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *model.ProductType) (*model.ProductType, error)
	GetAll() ([]*model.ProductType, error)
	GetById(id uint) (*model.ProductType, error)
	Update(model *model.ProductType) (*model.ProductType, error)
	Delete(id uint) error
	GetByUniqueFields(name string) (*model.ProductType, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *model.ProductType) (*model.ProductType, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll() ([]*model.ProductType, error) {
	var productTypes []*model.ProductType
	var err error

	result := r.DB.Find(&productTypes)
	if result.Error != nil {
		err = result.Error
	}

	return productTypes, err
}

func (r *repository) GetById(id uint) (*model.ProductType, error) {
	var productType *model.ProductType
	var err error

	result := r.DB.First(&productType, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return productType, err
}

func (r *repository) Update(model *model.ProductType) (*model.ProductType, error) {
	result := r.DB.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Delete(id uint) error {
	result := r.DB.Delete(&model.ProductType{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetByUniqueFields(name string) (*model.ProductType, error) {
	var productType *model.ProductType
	var err error

	result := r.DB.First(&productType, "name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return productType, err
}
