package category

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *Category) (*Category, error)
	GetAll() ([]*Category, error)
	GetById(id uint) (*Category, error)
	Update() (*Category, error)
	Delete(id int) error
	GetByUniqueFields(name, slug string) (*Category, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *Category) (*Category, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll() ([]*Category, error) {
	var categories []*Category
	var err error

	result := r.DB.Find(&categories)
	if result.Error != nil {
		err = result.Error
	}

	return categories, err
}

func (r *repository) GetById(id uint) (*Category, error) {
	var category *Category
	var err error

	result := r.DB.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return category, err
}

func (r *repository) Update() (*Category, error) {
	return nil, nil
}

func (r *repository) Delete(id int) error {
	return nil
}

func (r *repository) GetByUniqueFields(name, slug string) (*Category, error) {
	var category *Category
	var err error

	result := r.DB.First(&category, "name = ? AND slug = ?", name, slug)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return category, err
}
