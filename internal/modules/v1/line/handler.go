package line

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/line/dto"
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

// Create creates a new line
//
//	@Summary		Create a new line
//	@Description	Create a new line
//	@Tags			Line
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			line	body		dto.CreateDTO				true	"Line to create"
//	@Success		201		{object}	docsResponse.LineCreate201	"Line created successfully"
//	@Failure		400		{object}	docsResponse.LineCreate400	"Bad Request or Validation Error"
//	@Failure		401		{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403		{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500		{object}	docsResponse.Response500	"Server Error"
//	@Router			/api/v1/line/create [post]
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

	createdLine, errFields, err := h.svc.Create(createDto)
	if err != nil {
		msg := fmt.Sprintf("failed to create line: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusCreated, createdLine)
}

// GetAll retrieves all lines
//
//	@Summary		Get all lines
//	@Description	Retrieve all lines
//	@Tags			Line
//	@Security		AppAuth
//	@Produce		json
//	@Success		200	{object}	docsResponse.LineList200	"List of lines"
//	@Failure		403	{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500	{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/line [get]
func (h *Handler) GetAll(w http.ResponseWriter) {
	lines, err := h.svc.GetAll()
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve lines: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, lines)
}

// GetById retrieves a line by its ID
//
//	@Summary		Get line by ID
//	@Description	Retrieve line by its ID
//	@Tags			Line
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int							true	"Line ID"
//	@Success		200	{object}	docsResponse.LineGetById200	"Line found"
//	@Failure		400	{object}	docsResponse.Response400	"Invalid ID"
//	@Failure		403	{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404	"Line not found"
//	@Failure		500	{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/line/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing line id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid line id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	line, err := h.svc.GetById(uint(id))
	if line == nil {
		msg := fmt.Sprintf("line with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve line: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, line)
}

// Update updates a line by its ID
//
//	@Summary		Update line
//	@Description	Update line by ID
//	@Tags			Line
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Line ID"
//	@Param			line	body		dto.UpdateDTO				true	"Line update payload"
//	@Success		200		{object}	docsResponse.LineUpdate200	"Line updated"
//	@Failure		400		{object}	docsResponse.LineUpdate400	"Bad request or validation error"
//	@Failure		401		{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403		{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500		{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/line/{id}/update [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing line id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid line id: %s", idStr)
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

	updatedLine, errFields, err := h.svc.Update(uint(id), updateDto)
	if err != nil {
		msg := fmt.Sprintf("failed to update line: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusOK, updatedLine)
}

// Delete deletes a line by its ID
//
//	@Summary		Delete line
//	@Description	Delete line by ID
//	@Tags			Line
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int							true	"Line ID"
//	@Success		200	{object}	docsResponse.LineDelete200	"Line deleted"
//	@Failure		400	{object}	docsResponse.Response400	"Invalid ID"
//	@Failure		401	{object}	docsResponse.Response401	"Unauthorized"
//	@Failure		403	{object}	docsResponse.Response403	"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404	"Line not found"
//	@Failure		409	{object}	docsResponse.Response409	"Line has linked entities"
//	@Failure		500	{object}	docsResponse.Response500	"Server error"
//	@Router			/api/v1/line/{id}/delete [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing line id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid line id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	line, err := h.svc.Delete(uint(id))
	if line == nil {
		msg := fmt.Sprintf("line with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to delete line: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, line)
}
