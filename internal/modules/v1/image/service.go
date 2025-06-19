package image

import (
	"haircompany-shop-rest/internal/modules/v1/image/dto"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Service interface {
	UploadImage(file multipart.File, filename string) (*dto.ResponseDTO, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) UploadImage(file multipart.File, filename string) (*dto.ResponseDTO, error) {
	var imageDTO *dto.ResponseDTO
	uploadDir := "uploads/tmp"

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		log.Printf("failed to create upload directory: %v", err)
		return nil, err
	}

	tempFilePath := filepath.Join(uploadDir, filename)
	outFile, err := os.Create(tempFilePath)
	if err != nil {
		log.Printf("failed to create temporary file: %v", err)
		return nil, err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			log.Printf("failed to close temporary file: %v", err)
		}
	}(outFile)

	if _, err := io.Copy(outFile, file); err != nil {
		log.Printf("failed to write file contents: %v", err)
		return nil, err
	}

	imageDTO = TransformImageToResponseDTO(filename)

	return imageDTO, nil
}
