package constraint

import (
	"errors"
	"github.com/go-playground/validator/v10"
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

func ValidateImage(files []*multipart.FileHeader, imageType string) error {
	switch imageType {
	case "category":
		return validateCategoryImages(files)
	default:
		return errors.New("unsupported image type")
	}
}

func validateCategoryImages(files []*multipart.FileHeader) error {
	if len(files) == 0 {
		return errors.New("no images provided")
	}

	for _, file := range files {
		if file.Size > 1*1024*1024 { // 1 MB limit for category images
			return errors.New("image size exceeds 1 MB limit")
		}
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			return errors.New("invalid image type, only JPEG and PNG are allowed")
		}
	}

	return nil
}
