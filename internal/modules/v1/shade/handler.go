package shade

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/shade/dto"
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

// Create creates a new shade
//
//	@Summary		Create a new shade
//	@Description	Create a new shade
//	@Tags			Shade
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			shade	body		dto.CreateDTO				true	"Shade to create"
//	@Success		201		{object}	docsResponse.ShadeCreate201	"Shade created successfully"
//	@Failure		400		{object}	docsResponse.ShadeCreate400	"Bad Request or Validation Error"
//	@Failure		401		{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403		{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500		{object}	docsResponse.Response500	"Server Error"
//	@Router			/api/v1/shade/create [post]
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

	createdShade, err := h.svc.Create(createDto)
	if err != nil {
		msg := fmt.Sprintf("failed to create shade: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusCreated, createdShade)
}

// GetAll retrieves all shades
//
//	@Summary		Get all shades
//	@Description	Retrieve all shades
//	@Tags			Shade
//	@Security		AppAuth
//	@Produce		json
//	@Success		200	{object}	docsResponse.ShadeList200	"List of shades"
//	@Failure		403	{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500	{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/shade [get]
func (h *Handler) GetAll(w http.ResponseWriter) {
	shades, err := h.svc.GetAll()
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve shades: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, shades)
}

// GetById retrieves a shade by its ID
//
//	@Summary		Get shade by ID
//	@Description	Retrieve shade by its ID
//	@Tags			Shade
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int								true	"Shade ID"
//	@Success		200	{object}	docsResponse.ShadeGetById200	"Shade found"
//	@Failure		400	{object}	docsResponse.Response400		"Invalid ID"
//	@Failure		403	{object}	docsResponse.Response403		"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404		"Shade not found"
//	@Failure		500	{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/shade/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing shade id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid shade id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	shade, err := h.svc.GetById(uint(id))
	if shade == nil {
		msg := fmt.Sprintf("shade with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve shade: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, shade)
}

// Update updates a shade by its ID
//
//	@Summary		Update shade
//	@Description	Update shade by ID
//	@Tags			Shade
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Shade ID"
//	@Param			shade	body		dto.UpdateDTO				true	"Shade update payload"
//	@Success		200		{object}	docsResponse.ShadeUpdate200	"Shade updated"
//	@Failure		400		{object}	docsResponse.ShadeUpdate400	"Bad request or validation error"
//	@Failure		401		{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403		{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500		{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/shade/{id}/update [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing shade id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid shade id: %s", idStr)
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

	updatedShade, err := h.svc.Update(uint(id), updateDto)
	if err != nil {
		msg := fmt.Sprintf("failed to update shade: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, updatedShade)
}

// Delete deletes a shade by its ID
//
//	@Summary		Delete shade
//	@Description	Delete shade by ID
//	@Tags			Shade
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int							true	"Shade ID"
//	@Success		200	{object}	docsResponse.ShadeDelete200	"Shade deleted"
//	@Failure		400	{object}	docsResponse.Response400	"Invalid ID"
//	@Failure		401	{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403	{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404	"Shade not found"
//	@Failure		409	{object}	docsResponse.Response409	"Shade has linked entities"
//	@Failure		500	{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/shade/{id}/delete [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing shade id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid shade id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	shade, err := h.svc.Delete(uint(id))
	if shade == nil {
		msg := fmt.Sprintf("shade with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to delete shade: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, shade)
}
