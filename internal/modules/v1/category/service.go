package category

import (
	"haircompany-shop-rest/internal/modules/v1/category/dto"
	"haircompany-shop-rest/pkg/response"
	"log"
)

type Service interface {
	Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error)
	GetAll() ([]*dto.ResponseDTO, error)
	GetById(id uint) (*dto.ResponseDTO, error)
	Update() (*Category, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (c *service) Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	var validationErrors []response.ErrorField
	existingCategory, err := c.repo.GetByUniqueFields(createDto.Name, createDto.Slug)
	if err != nil {
		return nil, nil, err
	}
	if existingCategory != nil {
		if existingCategory.Name == createDto.Name {
			validationErrors = append(validationErrors, response.ErrorField{
				Field:     "name",
				ErrorCode: string(response.NotUnique),
			})
		}
		if existingCategory.Slug == createDto.Slug {
			validationErrors = append(validationErrors, response.ErrorField{
				Field:     "slug",
				ErrorCode: string(response.NotUnique),
			})
		}

		return nil, validationErrors, nil
	}

	if createDto.ParentID != nil {
		existingParent, err := c.repo.GetById(*createDto.ParentID)
		if err != nil || existingParent == nil {
			validationErrors = append(validationErrors, response.ErrorField{
				Field:     "parentId",
				ErrorCode: string(response.NotFound),
			})
			return nil, validationErrors, nil
		}
	}

	categoryModel := TransformCreateDTOToModel(createDto)
	createdCategory, err := c.repo.Create(categoryModel)
	if err != nil {
		return nil, nil, err
	}

	createdCategoryResponse := TransformModelToResponseDTO(createdCategory)

	return createdCategoryResponse, nil, nil
}

func (c *service) GetAll() ([]*dto.ResponseDTO, error) {
	categoryDTOs := make([]*dto.ResponseDTO, 0)
	models, err := c.repo.GetAll()
	if err != nil {
		log.Printf("error retrieving categories: %v", err)
	}

	for _, model := range models {
		categoryResponse := TransformModelToResponseDTO(model)
		categoryDTOs = append(categoryDTOs, categoryResponse)
	}

	return categoryDTOs, err
}

func (c *service) GetById(id uint) (*dto.ResponseDTO, error) {
	model, err := c.repo.GetById(id)
	if model == nil {
		return nil, err
	}

	categoryDTO := TransformModelToResponseDTO(model)

	return categoryDTO, err
}

func (c *service) Update() (*Category, error) {
	return c.repo.Update()
}

func (c *service) Delete(id int) error {
	return c.repo.Delete(id)
}
