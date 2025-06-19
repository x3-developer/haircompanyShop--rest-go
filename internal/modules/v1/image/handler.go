package image

import (
	"fmt"
	"haircompany-shop-rest/internal/modules/v1/image/dto"
	"haircompany-shop-rest/pkg/constraint"
	"haircompany-shop-rest/pkg/response"
	"mime/multipart"
	"net/http"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		svc: s,
	}
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		msg := fmt.Sprintf("failed to parse multipart form: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	form := r.MultipartForm
	imageType := form.Value["imageType"]
	if imageType == nil || len(imageType) == 0 || imageType[0] == "" {
		msg := "imageType is required"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}
	if err := constraint.ValidateImageType(imageType[0]); err != nil {
		msg := fmt.Sprintf("invalid imageType: %s", imageType[0])
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	files := form.File["images"]
	if err := constraint.ValidateImage(files, imageType[0]); err != nil {
		msg := fmt.Sprintf("image validation failed: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	var uploadedImages []*dto.ResponseDTO

	for _, fileHeader := range files {
		func(fileHeader *multipart.FileHeader) {
			file, err := fileHeader.Open()
			if err != nil {
				msg := fmt.Sprintf("failed to open file: %v", err)
				response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
				return
			}
			defer func(file multipart.File) {
				err := file.Close()
				if err != nil {
					msg := fmt.Sprintf("failed to close file: %v", err)
					response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
					return
				}
			}(file)

			imageDTO, err := h.svc.UploadImage(file, fileHeader.Filename)
			if err != nil {
				msg := fmt.Sprintf("failed to upload image: %v", err)
				response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
				return
			}

			uploadedImages = append(uploadedImages, imageDTO)
		}(fileHeader)
	}
	if len(uploadedImages) == 0 {
		msg := "no images were uploaded"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	response.SendSuccess(w, http.StatusOK, uploadedImages)
}
