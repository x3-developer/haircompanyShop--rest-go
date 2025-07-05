package product_type

import (
	"errors"
	"haircompany-shop-rest/internal/modules/v1/product_type/dto"
	"haircompany-shop-rest/pkg/response"
	"log"
)

type Service interface {
	Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error)
	GetAll() ([]*dto.ResponseDTO, error)
	GetById(id uint) (*dto.ResponseDTO, error)
	Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error)
	Delete(id uint) (*dto.ResponseDTO, error)
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
	existingProductType, err := c.repo.GetByUniqueFields(createDto.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingProductType != nil {
		if existingProductType.Name == createDto.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}

		return nil, validationErrors, nil
	}

	productTypeModel := dto.TransformCreateDTOToModel(createDto)
	createdProductType, err := c.repo.Create(productTypeModel)
	if err != nil {
		return nil, nil, err
	}

	createdProductTypeResponse := dto.TransformModelToResponseDTO(createdProductType)

	return createdProductTypeResponse, nil, nil
}

func (c *service) GetAll() ([]*dto.ResponseDTO, error) {
	productTypeDTOs := make([]*dto.ResponseDTO, 0)
	models, err := c.repo.GetAll()
	if err != nil {
		log.Printf("error retrieving product types: %v", err)
	}

	for _, model := range models {
		productTypeResponse := dto.TransformModelToResponseDTO(model)
		productTypeDTOs = append(productTypeDTOs, productTypeResponse)
	}

	return productTypeDTOs, err
}

func (c *service) GetById(id uint) (*dto.ResponseDTO, error) {
	model, err := c.repo.GetById(id)
	if model == nil {
		return nil, err
	}

	productTypeDTO := dto.TransformModelToResponseDTO(model)

	return productTypeDTO, err
}

func (c *service) Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	var validationErrors []response.ErrorField
	model, err := c.repo.GetById(id)
	if err != nil {
		return nil, nil, err

	}
	if model == nil {
		return nil, nil, errors.New("product type not found")
	}

	dto.TransformUpdateDTOToModel(updateDto, model)
	existingProductType, err := c.repo.GetByUniqueFields(model.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingProductType != nil && existingProductType.ID != id {
		if existingProductType.Name == model.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}
		return nil, validationErrors, nil
	}

	updatedProductType, err := c.repo.Update(model)
	if err != nil {
		return nil, nil, err
	}

	updatedProductTypeResponse := dto.TransformModelToResponseDTO(updatedProductType)

	return updatedProductTypeResponse, nil, nil

}

func (c *service) Delete(id uint) (*dto.ResponseDTO, error) {
	existedProductType, err := c.repo.GetById(id)
	if existedProductType == nil {
		return nil, err
	}

	productTypeDTO := dto.TransformModelToResponseDTO(existedProductType)

	err = c.repo.Delete(id)
	if err != nil {
		return productTypeDTO, err
	}

	return productTypeDTO, nil
}
