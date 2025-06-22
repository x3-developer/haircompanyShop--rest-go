package image

import (
	"haircompany-shop-rest/internal/modules/v1/image/dto"
	"haircompany-shop-rest/internal/services"
	"log"
	"mime/multipart"
)

type Service interface {
	UploadImage(file multipart.File, filename string) (*dto.ResponseDTO, error)
}

type service struct {
	fileService services.FileSystemService
}

func NewService(fs services.FileSystemService) Service {
	return &service{
		fileService: fs,
	}
}

func (s *service) UploadImage(file multipart.File, filename string) (*dto.ResponseDTO, error) {
	var imageDTO *dto.ResponseDTO
	newFilename, err := s.fileService.SaveToTemp(file, filename)
	if err != nil {
		log.Printf("failed to upload image: %v", err)
		return nil, err
	}

	imageDTO = TransformImageToResponseDTO(newFilename)

	return imageDTO, nil
}
