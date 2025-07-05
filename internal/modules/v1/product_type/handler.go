package product_type

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/product_type/dto"
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

// Create creates a new productType
//
//	@Summary		Create a new productType
//	@Description	Create a new productType
//	@Tags			ProductType
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			productType	body		dto.CreateDTO						true	"ProductType to create"
//	@Success		201			{object}	docsResponse.ProductTypeCreate201	"ProductType created successfully"
//	@Failure		400			{object}	docsResponse.ProductTypeCreate400	"Bad Request or Validation Error"
//	@Failure		401			{object}	docsResponse.Response401			"Unauthorized"
//	@Failure		403			{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500			{object}	docsResponse.Response500			"Server Error"
//	@Router			/api/v1/product-type/create [post]
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

	createdProductType, errFields, err := h.svc.Create(createDto)
	if err != nil {
		msg := fmt.Sprintf("failed to create product type: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusCreated, createdProductType)
}

// GetAll retrieves all productTypes
//
//	@Summary		Get all productTypes
//	@Description	Retrieve all productTypes
//	@Tags			ProductType
//	@Security		AppAuth
//	@Produce		json
//	@Success		200	{object}	docsResponse.ProductTypeList200	"List of productTypes"
//	@Failure		403	{object}	docsResponse.Response403		"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500	{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/product-type [get]
func (h *Handler) GetAll(w http.ResponseWriter) {
	productTypes, err := h.svc.GetAll()
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve product types: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, productTypes)
}

// GetById retrieves a productType by its ID
//
//	@Summary		Get productType by ID
//	@Description	Retrieve productType by its ID
//	@Tags			ProductType
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int									true	"ProductType ID"
//	@Success		200	{object}	docsResponse.ProductTypeGetById200	"ProductType found"
//	@Failure		400	{object}	docsResponse.Response400			"Invalid ID"
//	@Failure		403	{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404			"ProductType not found"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/product-type/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing product type id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid product yype id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	productType, err := h.svc.GetById(uint(id))
	if productType == nil {
		msg := fmt.Sprintf("product type with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve product type: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, productType)
}

// Update updates a productType by its ID
//
//	@Summary		Update productType
//	@Description	Update productType by ID
//	@Tags			ProductType
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int									true	"ProductType ID"
//	@Param			productType	body		dto.UpdateDTO						true	"ProductType update payload"
//	@Success		200			{object}	docsResponse.ProductTypeUpdate200	"ProductType updated"
//	@Failure		400			{object}	docsResponse.ProductTypeUpdate400	"Bad request or validation error"
//	@Failure		401			{object}	docsResponse.Response401			"Unauthorized"
//	@Failure		403			{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500			{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/product-type/{id}/update [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing product type id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid product type id: %s", idStr)
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

	updatedProductType, errFields, err := h.svc.Update(uint(id), updateDto)
	if err != nil {
		msg := fmt.Sprintf("failed to update product type: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusOK, updatedProductType)
}

// Delete deletes a productType by its ID
//
//	@Summary		Delete productType
//	@Description	Delete productType by ID
//	@Tags			ProductType
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int									true	"ProductType ID"
//	@Success		200	{object}	docsResponse.ProductTypeDelete200	"ProductType deleted"
//	@Failure		400	{object}	docsResponse.Response400			"Invalid ID"
//	@Failure		401	{object}	docsResponse.Response401			"Unauthorized"
//	@Failure		403	{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404			"ProductType not found"
//	@Failure		409	{object}	docsResponse.Response409			"ProductType has linked entities"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/product-type/{id}/delete [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing product type id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid product type id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	productType, err := h.svc.Delete(uint(id))
	if productType == nil {
		msg := fmt.Sprintf("product type with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to delete product type: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, productType)
}
