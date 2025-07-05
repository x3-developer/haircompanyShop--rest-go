package line

import (
	"errors"
	"gorm.io/gorm"
	"haircompany-shop-rest/internal/modules/v1/line/model"
	"haircompany-shop-rest/pkg/database"
)

type Repository interface {
	Create(model *model.Line) (*model.Line, error)
	GetAll() ([]*model.Line, error)
	GetById(id uint) (*model.Line, error)
	Update(model *model.Line) (*model.Line, error)
	Delete(id uint) error
	GetByUniqueFields(name string) (*model.Line, error)
}

type repository struct {
	DB *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(model *model.Line) (*model.Line, error) {
	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) GetAll() ([]*model.Line, error) {
	var lines []*model.Line
	var err error

	result := r.DB.Find(&lines)
	if result.Error != nil {
		err = result.Error
	}

	return lines, err
}

func (r *repository) GetById(id uint) (*model.Line, error) {
	var line *model.Line
	var err error

	result := r.DB.First(&line, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return line, err
}

func (r *repository) Update(model *model.Line) (*model.Line, error) {
	result := r.DB.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return model, nil
}

func (r *repository) Delete(id uint) error {
	result := r.DB.Delete(&model.Line{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetByUniqueFields(name string) (*model.Line, error) {
	var line *model.Line
	var err error

	result := r.DB.First(&line, "name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = result.Error
	}

	return line, err
}
