package category

import (
	"context"
	"errors"
	"haircompany-shop-rest/internal/modules/v1/category/dto"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/response"
	"haircompany-shop-rest/pkg/utils"
	"log"
	"sync"
)

type Service interface {
	Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error)
	GetAll() ([]*dto.ResponseDTO, error)
	GetById(id uint) (*dto.ResponseDTO, error)
	Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error)
	Delete(id uint) error
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

func (c *service) Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	var validationErrors []response.ErrorField
	existingCategory, err := c.repo.GetByUniqueFields(createDto.Name, createDto.Slug)
	if err != nil {
		return nil, nil, err
	}
	if existingCategory != nil {
		if existingCategory.Name == createDto.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}
		if existingCategory.Slug == createDto.Slug {
			validationErrors = append(validationErrors, response.NewErrorField("slug", string(response.NotUnique)))
		}

		return nil, validationErrors, nil
	}

	if createDto.ParentID != nil {
		existingParent, err := c.repo.GetById(*createDto.ParentID)
		if err != nil || existingParent == nil {
			validationErrors = append(validationErrors, response.NewErrorField("parentId", string(response.NotFound)))
			return nil, validationErrors, nil
		}
	}

	categoryModel := TransformCreateDTOToModel(createDto)
	createdCategory, err := c.repo.Create(categoryModel)
	if err != nil {
		return nil, nil, err
	}

	filenames := []string{categoryModel.Image, categoryModel.HeaderImage}
	utils.SafeGo(c.ctx, c.wg, "MoveImageToPermanent", func(ctx context.Context) {
		if ctx.Err() != nil {
			log.Println("context cancelled, skipping image move")
			return
		}

		if err := c.fileService.MoveToPermanent(filenames, "images/category"); err != nil {
			log.Printf("error moving images to permanent storage: %v", err)
		}
	})

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

func (c *service) Update(id uint, updateDto dto.UpdateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	model, err := c.repo.GetById(id)
	if err != nil {
		return nil, nil, err

	}
	if model == nil {
		return nil, nil, errors.New("category not found")
	}

	TransformUpdateDTOToModel(updateDto, model)
	existingCategory, err := c.repo.GetByUniqueFields(model.Name, model.Slug)
	if err != nil {
		return nil, nil, err
	}
	if existingCategory != nil && existingCategory.ID != id {
		var validationErrors []response.ErrorField
		if existingCategory.Name == model.Name {
			validationErrors = append(validationErrors, response.NewErrorField("name", string(response.NotUnique)))
		}
		if existingCategory.Slug == model.Slug {
			validationErrors = append(validationErrors, response.NewErrorField("slug", string(response.NotUnique)))
		}
		return nil, validationErrors, nil
	}

	updatedCategory, err := c.repo.Update(model)
	if err != nil {
		return nil, nil, err
	}

	var filenames []string

	if updateDto.Image != nil {
		filenames = append(filenames, *updateDto.Image)
	}
	if updateDto.HeaderImage != nil {
		filenames = append(filenames, *updateDto.HeaderImage)
	}

	if len(filenames) != 0 {
		utils.SafeGo(c.ctx, c.wg, "MoveImageToPermanent", func(ctx context.Context) {
			if ctx.Err() != nil {
				log.Println("context cancelled, skipping image move")
				return
			}

			if err := c.fileService.MoveToPermanent(filenames, "images/category"); err != nil {
				log.Printf("error moving images to permanent storage: %v", err)
			}
		})
	}

	updatedCategoryResponse := TransformModelToResponseDTO(updatedCategory)

	return updatedCategoryResponse, nil, nil

}

func (c *service) Delete(id uint) error {
	return c.repo.Delete(id)
}
