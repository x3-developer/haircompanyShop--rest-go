package image

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"haircompany-shop-rest/internal/modules/v1/image/dto"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const uploadDir = "uploads"

type Service interface {
	UploadImageToTemp(file multipart.File, filename string) (*dto.ResponseDTO, error)
	MoveImageToPermanent(filenames []string, entityName string)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) UploadImageToTemp(file multipart.File, filename string) (*dto.ResponseDTO, error) {
	var imageDTO *dto.ResponseDTO
	uploadDir := filepath.Join(uploadDir, "temp")

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		log.Printf("failed to create upload directory: %v", err)
		return nil, err
	}

	filename, err = getNewImageName(filename)
	if err != nil {
		log.Printf("failed to generate new image name: %v", err)
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

func (s *service) MoveImageToPermanent(filenames []string, entityName string) {
	permanentDir, err := getPermanentImagePath(entityName)
	if err != nil {
		log.Printf("failed to get permanent image path: %v", err)
		return
	}

	if err := os.MkdirAll(permanentDir, os.ModePerm); err != nil {
		log.Printf("failed to create permanent upload directory: %v", err)
		return
	}

	for _, filename := range filenames {
		tempFilePath := filepath.Join(uploadDir, "temp", filename)
		permanentFilePath := filepath.Join(permanentDir, filename)
		if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
			log.Printf("temporary file does not exist: %s", tempFilePath)
			continue
		}

		if err := os.Rename(tempFilePath, permanentFilePath); err != nil {
			log.Printf("failed to move file from %s to %s: %v", tempFilePath, permanentFilePath, err)
			continue
		}
	}
}

func getNewImageName(filename string) (string, error) {
	randomBytes := make([]byte, 8)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	randomPrefix := hex.EncodeToString(randomBytes)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	newName := name + "_" + randomPrefix + ext

	return newName, nil
}

func getPermanentImagePath(entityName string) (string, error) {
	if entityName == "" {
		return "", errors.New("entity name cannot be empty")
	}

	uploadDir := filepath.Join(uploadDir, "images")
	uploadDir = filepath.Join(uploadDir, entityName)

	return uploadDir, nil
}
