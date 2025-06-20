package constraint

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"haircompany-shop-rest/pkg/response"
	"mime/multipart"
)

func ValidateDTO(dto interface{}) error {
	validate := validator.New()

	return validate.Struct(dto)
}

func ValidateImageType(t string) error {
	switch t {
	case "category":
		return nil
	default:
		return errors.New("unsupported image type")
	}
}

func ValidateImage(files []*multipart.FileHeader, imageType string) ([]response.ErrorField, error) {
	switch imageType {
	case "category":
		return validateCategoryImages(files)
	default:
		return nil, errors.New("unsupported image type")
	}
}

func validateCategoryImages(files []*multipart.FileHeader) ([]response.ErrorField, error) {
	if len(files) == 0 {
		return nil, errors.New("no images provided")
	}

	var validationErrors []response.ErrorField

	for i, file := range files {
		if file.Size > 1*1024*1024 { // 1 MB limit for category images
			validationErrors = append(validationErrors, response.NewErrorField(fmt.Sprintf("image_%d", i), string(response.FileTooLarge)))
		}
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			validationErrors = append(validationErrors, response.NewErrorField(fmt.Sprintf("image_%d", i), string(response.InvalidFileType)))
		}
	}

	return validationErrors, nil
}
