package shade

import (
	"context"
	"errors"
	"haircompany-shop-rest/internal/modules/v1/shade/dto"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/utils"
	"log"
	"sync"
)

type Service interface {
	Create(createDto dto.CreateDTO) (*dto.ResponseDTO, error)
	GetAll() ([]*dto.ResponseDTO, error)
	GetById(id uint) (*dto.ResponseDTO, error)
	Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, error)
	Delete(id uint) (*dto.ResponseDTO, error)
}

type service struct {
	repo        Repository
	fileService services.FileSystemService
	ctx         context.Context
	wg          *sync.WaitGroup
}

func NewService(r Repository, fs services.FileSystemService, ctx context.Context, wg *sync.WaitGroup) Service {
	return &service{
		repo:        r,
		fileService: fs,
		ctx:         ctx,
		wg:          wg,
	}
}

func (c *service) Create(createDto dto.CreateDTO) (*dto.ResponseDTO, error) {
	shadeModel := dto.TransformCreateDTOToModel(createDto)
	createdShade, err := c.repo.Create(shadeModel)
	if err != nil {
		return nil, err
	}

	filenames := []string{shadeModel.Image}
	utils.SafeGo(c.ctx, c.wg, "MoveImageToPermanent", func(ctx context.Context) {
		if ctx.Err() != nil {
			log.Println("context cancelled, skipping image move")
			return
		}

		if err := c.fileService.MoveToPermanent(filenames, "images/shade"); err != nil {
			log.Printf("error moving images to permanent storage: %v", err)
		}
	})

	createdShadeResponse := dto.TransformModelToResponseDTO(createdShade)

	return createdShadeResponse, nil
}

func (c *service) GetAll() ([]*dto.ResponseDTO, error) {
	shadeDTOs := make([]*dto.ResponseDTO, 0)
	models, err := c.repo.GetAll()
	if err != nil {
		log.Printf("error retrieving shades: %v", err)
	}

	for _, model := range models {
		shadeResponse := dto.TransformModelToResponseDTO(model)
		shadeDTOs = append(shadeDTOs, shadeResponse)
	}

	return shadeDTOs, err
}

func (c *service) GetById(id uint) (*dto.ResponseDTO, error) {
	model, err := c.repo.GetById(id)
	if model == nil {
		return nil, err
	}

	shadeDTO := dto.TransformModelToResponseDTO(model)

	return shadeDTO, err
}

func (c *service) Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, error) {
	model, err := c.repo.GetById(id)
	if err != nil {
		return nil, err

	}
	if model == nil {
		return nil, errors.New("shade not found")
	}

	dto.TransformUpdateDTOToModel(updateDto, model)

	updatedShade, err := c.repo.Update(model)
	if err != nil {
		return nil, err
	}

	var filenames []string

	if updateDto.Image != nil {
		filenames = append(filenames, *updateDto.Image)
	}

	if len(filenames) != 0 {
		utils.SafeGo(c.ctx, c.wg, "MoveImageToPermanent", func(ctx context.Context) {
			if ctx.Err() != nil {
				log.Println("context cancelled, skipping image move")
				return
			}

			if err := c.fileService.MoveToPermanent(filenames, "images/shade"); err != nil {
				log.Printf("error moving images to permanent storage: %v", err)
			}
		})
	}

	updatedShadeResponse := dto.TransformModelToResponseDTO(updatedShade)

	return updatedShadeResponse, nil
}

func (c *service) Delete(id uint) (*dto.ResponseDTO, error) {
	existedShade, err := c.repo.GetById(id)
	if existedShade == nil {
		return nil, err
	}

	shadeDTO := dto.TransformModelToResponseDTO(existedShade)

	err = c.repo.Delete(id)
	if err != nil {
		return shadeDTO, err
	}

	return shadeDTO, nil
}
