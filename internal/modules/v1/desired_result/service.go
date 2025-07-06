package desired_result

import (
	"errors"
	"haircompany-shop-rest/internal/modules/v1/desired_result/dto"
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
	existingDesiredResult, err := c.repo.GetByUniqueFields(createDto.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingDesiredResult != nil {
		if existingDesiredResult.Name == createDto.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}

		return nil, validationErrors, nil
	}

	desiredResultModel := dto.TransformCreateDTOToModel(createDto)
	createdDesiredResult, err := c.repo.Create(desiredResultModel)
	if err != nil {
		return nil, nil, err
	}

	createdDesiredResultResponse := dto.TransformModelToResponseDTO(createdDesiredResult)

	return createdDesiredResultResponse, nil, nil
}

func (c *service) GetAll() ([]*dto.ResponseDTO, error) {
	desiredResultDTOs := make([]*dto.ResponseDTO, 0)
	models, err := c.repo.GetAll()
	if err != nil {
		log.Printf("error retrieving desired results: %v", err)
	}

	for _, model := range models {
		desiredResultResponse := dto.TransformModelToResponseDTO(model)
		desiredResultDTOs = append(desiredResultDTOs, desiredResultResponse)
	}

	return desiredResultDTOs, err
}

func (c *service) GetById(id uint) (*dto.ResponseDTO, error) {
	model, err := c.repo.GetById(id)
	if model == nil {
		return nil, err
	}

	desiredResultDTO := dto.TransformModelToResponseDTO(model)

	return desiredResultDTO, err
}

func (c *service) Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	var validationErrors []response.ErrorField
	model, err := c.repo.GetById(id)
	if err != nil {
		return nil, nil, err

	}
	if model == nil {
		return nil, nil, errors.New("desired result not found")
	}

	dto.TransformUpdateDTOToModel(updateDto, model)
	existingDesiredResult, err := c.repo.GetByUniqueFields(model.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingDesiredResult != nil && existingDesiredResult.ID != id {
		if existingDesiredResult.Name == model.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}
		return nil, validationErrors, nil
	}

	updatedDesiredResult, err := c.repo.Update(model)
	if err != nil {
		return nil, nil, err
	}

	updatedDesiredResultResponse := dto.TransformModelToResponseDTO(updatedDesiredResult)

	return updatedDesiredResultResponse, nil, nil

}

func (c *service) Delete(id uint) (*dto.ResponseDTO, error) {
	existedDesiredResult, err := c.repo.GetById(id)
	if existedDesiredResult == nil {
		return nil, err
	}

	desiredResultDTO := dto.TransformModelToResponseDTO(existedDesiredResult)

	err = c.repo.Delete(id)
	if err != nil {
		return desiredResultDTO, err
	}

	return desiredResultDTO, nil
}
