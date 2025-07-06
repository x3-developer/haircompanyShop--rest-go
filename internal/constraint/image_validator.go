package constraint

import (
	"errors"
	"fmt"
	"haircompany-shop-rest/pkg/response"
	"mime/multipart"
)

var imageTypes = []string{"category", "shade"}

func ValidateImageType(t string) error {
	for _, validType := range imageTypes {
		if t == validType {
			return nil
		}
	}
	return fmt.Errorf("invalid image type: %s", t)
}

func ValidateImage(files []*multipart.FileHeader, imageType string) ([]response.ErrorField, error) {
	if err := ValidateImageType(imageType); err != nil {
		return nil, err
	}

	validationErrors, err := validateImages(files, imageType)
	if err != nil {
		return nil, err
	}

	if len(validationErrors) > 0 {
		return validationErrors, nil
	}

	return nil, nil
}

func validateImages(files []*multipart.FileHeader, imageType string) ([]response.ErrorField, error) {
	if len(files) == 0 {
		return nil, errors.New("no images provided")
	}

	var validationErrors []response.ErrorField
	var imageSizeLimit int64
	var imageContentType []string

	switch imageType {
	case "category":
		imageSizeLimit = 1 * 1024 * 1024 // 1 MB limit for category images
		imageContentType = []string{"image/jpeg", "image/png"}
		break
	case "shade":
		imageSizeLimit = 500 * 1024 // 500 KB limit for shade images
		imageContentType = []string{"image/jpeg", "image/png"}
		break
	default:
		return nil, fmt.Errorf("unsupported image type: %s", imageType)
	}

	for i, file := range files {
		if file.Size > imageSizeLimit {
			validationErrors = append(validationErrors, response.NewErrorField(fmt.Sprintf("image_%d", i), string(response.FileTooLarge)))
		}
		fileHeader := file.Header.Get("Content-Type")

		validContentType := false
		for _, contentType := range imageContentType {
			if fileHeader == contentType {
				validContentType = true
				break
			}
		}
		if !validContentType {
			validationErrors = append(validationErrors, response.NewErrorField(fmt.Sprintf("image_%d", i), string(response.InvalidFileType)))
		}
	}

	return validationErrors, nil
}
