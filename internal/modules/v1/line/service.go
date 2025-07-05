package line

import (
	"errors"
	"haircompany-shop-rest/internal/modules/v1/line/dto"
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
	existingLine, err := c.repo.GetByUniqueFields(createDto.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingLine != nil {
		if existingLine.Name == createDto.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}

		return nil, validationErrors, nil
	}

	lineModel := dto.TransformCreateDTOToModel(createDto)
	createdLine, err := c.repo.Create(lineModel)
	if err != nil {
		return nil, nil, err
	}

	createdLineResponse := dto.TransformModelToResponseDTO(createdLine)

	return createdLineResponse, nil, nil
}

func (c *service) GetAll() ([]*dto.ResponseDTO, error) {
	lineDTOs := make([]*dto.ResponseDTO, 0)
	models, err := c.repo.GetAll()
	if err != nil {
		log.Printf("error retrieving lines: %v", err)
	}

	for _, model := range models {
		lineResponse := dto.TransformModelToResponseDTO(model)
		lineDTOs = append(lineDTOs, lineResponse)
	}

	return lineDTOs, err
}

func (c *service) GetById(id uint) (*dto.ResponseDTO, error) {
	model, err := c.repo.GetById(id)
	if model == nil {
		return nil, err
	}

	lineDTO := dto.TransformModelToResponseDTO(model)

	return lineDTO, err
}

func (c *service) Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	var validationErrors []response.ErrorField
	model, err := c.repo.GetById(id)
	if err != nil {
		return nil, nil, err

	}
	if model == nil {
		return nil, nil, errors.New("line not found")
	}

	dto.TransformUpdateDTOToModel(updateDto, model)
	existingLine, err := c.repo.GetByUniqueFields(model.Name)
	if err != nil {
		return nil, nil, err
	}
	if existingLine != nil && existingLine.ID != id {
		if existingLine.Name == model.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}
		return nil, validationErrors, nil
	}

	updatedLine, err := c.repo.Update(model)
	if err != nil {
		return nil, nil, err
	}

	updatedLineResponse := dto.TransformModelToResponseDTO(updatedLine)

	return updatedLineResponse, nil, nil

}

func (c *service) Delete(id uint) (*dto.ResponseDTO, error) {
	existedLine, err := c.repo.GetById(id)
	if existedLine == nil {
		return nil, err
	}

	lineDTO := dto.TransformModelToResponseDTO(existedLine)

	err = c.repo.Delete(id)
	if err != nil {
		return lineDTO, err
	}

	return lineDTO, nil
}
