package image

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/image/dto"
	"haircompany-shop-rest/pkg/response"
	"mime/multipart"
	"net/http"
	"strings"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		svc: s,
	}
}

// Upload images
//
//	@Summary		Upload images
//	@Description	Upload one or more images with a specific imageType. Requires JWT authentication.
//	@Tags			Image
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			imageType	formData	string						true	"Type of image (e.g. category)"
//	@Param			images		formData	file						true	"Image files to upload (multiple allowed)"
//	@Success		200			{array}		docsResponse.ImageUpload200	"Uploaded images info"
//	@Failure		400			{object}	docsResponse.ImageUpload400	"Bad Request or Validation Error"
//	@Failure		401			{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403			{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		413			{object}	docsResponse.Response413	"File too large"
//	@Failure		500			{object}	docsResponse.Response500	"Server Error"
//	@Router			/api/v1/image/upload [post]
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	bodyLimit := int64(50 << 20) // Limit request body to 50 MB
	r.Body = http.MaxBytesReader(w, r.Body, bodyLimit)
	if err := r.ParseMultipartForm(bodyLimit); err != nil {
		if strings.Contains(err.Error(), "http: request body too large") {
			response.SendError(w, http.StatusRequestEntityTooLarge, "file too large", response.RequestTooLarge)
			return
		}
		response.SendError(w, http.StatusBadRequest, fmt.Sprintf("failed to parse form: %v", err), response.BadRequest)
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
	errFields, err := constraint.ValidateImage(files, imageType[0])
	if err != nil {
		msg := fmt.Sprintf("image validation failed: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
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
