package category

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/category/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *model.Category) (*model.Category, error)
	GetAll() ([]*model.Category, error)
	GetById(id uint) (*model.Category, error)
	Update(model *model.Category) (*model.Category, error)
	Delete(id uint) error
	GetByUniqueFields(name, slug string) (*model.Category, error)
	CountChildrenByParentId(parentId uint) (int64, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *model.Category) (*model.Category, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll() ([]*model.Category, error) {
	var categories []*model.Category
	var err error

	result := r.DB.Find(&categories)
	if result.Error != nil {
		err = result.Error
	}

	return categories, err
}

func (r *repository) GetById(id uint) (*model.Category, error) {
	var category *model.Category
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

func (r *repository) Update(model *model.Category) (*model.Category, error) {
	result := r.DB.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Delete(id uint) error {
	result := r.DB.Delete(&model.Category{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetByUniqueFields(name, slug string) (*model.Category, error) {
	var category *model.Category
	var err error

	result := r.DB.First(&category, "name = ? OR slug = ?", name, slug)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return category, err
}

func (r *repository) CountChildrenByParentId(parentId uint) (int64, error) {
	var count int64
	result := r.DB.Model(&model.Category{}).Where("parent_id = ?", parentId).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}
