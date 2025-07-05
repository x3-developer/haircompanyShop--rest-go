package category

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/category/dto"
	"haircompany-shop-rest/pkg/request"
	"haircompany-shop-rest/pkg/response"
	"net/http"
	"strconv"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		svc: s,
	}
}

// Create creates a new category
//
//	@Summary		Create a new category
//	@Description	Create a new category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			category	body		dto.CreateDTO					true	"Category to create"
//	@Success		201			{object}	docsResponse.CategoryCreate201	"Category created successfully"
//	@Failure		400			{object}	docsResponse.CategoryCreate400	"Bad Request or Validation Error"
//	@Failure		500			{object}	docsResponse.Response500		"Server Error"
//	@Router			/api/v1/category [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	createDto, err := request.DecodeBody[dto.CreateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := constraint.ValidateDTO(createDto)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	createdCategory, errFields, err := h.svc.Create(createDto)
	if err != nil {
		msg := fmt.Sprintf("failed to create category: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusCreated, createdCategory)
}

// GetAll retrieves all categories
//
//	@Summary		Get all categories
//	@Description	Retrieve all categories
//	@Tags			Category
//	@Produce		json
//	@Success		200	{object}	docsResponse.CategoryList200	"List of categories"
//	@Failure		500	{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/category [get]
func (h *Handler) GetAll(w http.ResponseWriter) {
	categories, err := h.svc.GetAll()
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve categories: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, categories)
}

// GetById retrieves a category by its ID
//
//	@Summary		Get category by ID
//	@Description	Retrieve category by its ID
//	@Tags			Category
//	@Produce		json
//	@Param			id	path		int								true	"Category ID"
//	@Success		200	{object}	docsResponse.CategoryGetById200	"Category found"
//	@Failure		400	{object}	docsResponse.Response400		"Invalid ID"
//	@Failure		404	{object}	docsResponse.Response404		"Category not found"
//	@Failure		500	{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/category/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing category id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid category id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	category, err := h.svc.GetById(uint(id))
	if category == nil {
		msg := fmt.Sprintf("category with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve category: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, category)
}

// Update updates a category by its ID
//
//	@Summary		Update category
//	@Description	Update category by ID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int								true	"Category ID"
//	@Param			category	body		dto.UpdateDTO					true	"Category update payload"
//	@Success		200			{object}	docsResponse.CategoryUpdate200	"Category updated"
//	@Failure		400			{object}	docsResponse.CategoryUpdate400	"Bad request or validation error"
//	@Failure		500			{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/category/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing category id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid category id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	updateDto, err := request.DecodeBody[dto.UpdateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := constraint.ValidateDTO(updateDto)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	updatedCategory, errFields, err := h.svc.Update(uint(id), updateDto)
	if err != nil {
		msg := fmt.Sprintf("failed to update category: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusOK, updatedCategory)
}

// Delete deletes a category by its ID
//
//	@Summary		Delete category
//	@Description	Delete category by ID
//	@Tags			Category
//	@Produce		json
//	@Param			id	path		int								true	"Category ID"
//	@Success		200	{object}	docsResponse.CategoryDelete200	"Category deleted"
//	@Failure		400	{object}	docsResponse.Response400		"Invalid ID"
//	@Failure		404	{object}	docsResponse.Response404		"Category not found"
//	@Failure		409	{object}	docsResponse.Response409		"Category has linked entities"
//	@Failure		500	{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/category/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing category id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid category id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	category, linkedEntitiesCount, err := h.svc.Delete(uint(id))
	if category == nil {
		msg := fmt.Sprintf("category with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if linkedEntitiesCount > 0 {
		msg := fmt.Sprintf("category with id %d cannot be deleted because it has %d linked entities", id, linkedEntitiesCount)
		response.SendError(w, http.StatusConflict, msg, response.HasLinkedEntities)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to delete category: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, category)
}
