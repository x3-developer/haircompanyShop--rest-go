package desired_result

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/desired_result/dto"
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

// Create creates a new desiredResult
//
//	@Summary		Create a new desiredResult
//	@Description	Create a new desiredResult
//	@Tags			DesiredResult
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			desiredResult	body		dto.CreateDTO						true	"DesiredResult to create"
//	@Success		201			{object}	docsResponse.DesiredResultCreate201	"DesiredResult created successfully"
//	@Failure		400			{object}	docsResponse.DesiredResultCreate400	"Bad Request or Validation Error"
//	@Failure		401			{object}	docsResponse.Response401			"Unauthorized"
//	@Failure		403			{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500			{object}	docsResponse.Response500			"Server Error"
//	@Router			/api/v1/desired-result/create [post]
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

	createdDesiredResult, errFields, err := h.svc.Create(createDto)
	if err != nil {
		msg := fmt.Sprintf("failed to create desired result: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusCreated, createdDesiredResult)
}

// GetAll retrieves all desiredResults
//
//	@Summary		Get all desiredResults
//	@Description	Retrieve all desiredResults
//	@Tags			DesiredResult
//	@Security		AppAuth
//	@Produce		json
//	@Success		200	{object}	docsResponse.DesiredResultList200	"List of desiredResults"
//	@Failure		403	{object}	docsResponse.Response403		"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500	{object}	docsResponse.Response500		"Server error"
//	@Router			/api/v1/desired-result [get]
func (h *Handler) GetAll(w http.ResponseWriter) {
	desiredResults, err := h.svc.GetAll()
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve desired results: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, desiredResults)
}

// GetById retrieves a desiredResult by its ID
//
//	@Summary		Get desiredResult by ID
//	@Description	Retrieve desiredResult by its ID
//	@Tags			DesiredResult
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int									true	"DesiredResult ID"
//	@Success		200	{object}	docsResponse.DesiredResultGetById200	"DesiredResult found"
//	@Failure		400	{object}	docsResponse.Response400			"Invalid ID"
//	@Failure		403	{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404			"DesiredResult not found"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/desired-result/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing desired result id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid desired result id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	desiredResult, err := h.svc.GetById(uint(id))
	if desiredResult == nil {
		msg := fmt.Sprintf("desired result with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve desired result: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, desiredResult)
}

// Update updates a desiredResult by its ID
//
//	@Summary		Update desiredResult
//	@Description	Update desiredResult by ID
//	@Tags			DesiredResult
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int									true	"DesiredResult ID"
//	@Param			desiredResult	body		dto.UpdateDTO						true	"DesiredResult update payload"
//	@Success		200			{object}	docsResponse.DesiredResultUpdate200	"DesiredResult updated"
//	@Failure		400			{object}	docsResponse.DesiredResultUpdate400	"Bad request or validation error"
//	@Failure		401			{object}	docsResponse.Response401			"Unauthorized"
//	@Failure		403			{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		500			{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/desired-result/{id}/update [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing desired result id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid desired result id: %s", idStr)
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

	updatedDesiredResult, errFields, err := h.svc.Update(uint(id), updateDto)
	if err != nil {
		msg := fmt.Sprintf("failed to update desired result: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusOK, updatedDesiredResult)
}

// Delete deletes a desiredResult by its ID
//
//	@Summary		Delete desiredResult
//	@Description	Delete desiredResult by ID
//	@Tags			DesiredResult
//	@Security		BearerAuth
//	@Security		AppAuth
//	@Produce		json
//	@Param			id	path		int									true	"DesiredResult ID"
//	@Success		200	{object}	docsResponse.DesiredResultDelete200	"DesiredResult deleted"
//	@Failure		400	{object}	docsResponse.Response400			"Invalid ID"
//	@Failure		401	{object}	docsResponse.Response401			"Unauthorized"
//	@Failure		403	{object}	docsResponse.Response403			"Forbidden - Invalid X-AUTH-APP"
//	@Failure		404	{object}	docsResponse.Response404			"DesiredResult not found"
//	@Failure		409	{object}	docsResponse.Response409			"DesiredResult has linked entities"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/desired-result/{id}/delete [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing desired result id"
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid desired result id: %s", idStr)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	desiredResult, err := h.svc.Delete(uint(id))
	if desiredResult == nil {
		msg := fmt.Sprintf("desired result with id %d not found", id)
		response.SendError(w, http.StatusNotFound, msg, response.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to delete desired result: %v", err)
		response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
		return
	}

	response.SendSuccess(w, http.StatusOK, desiredResult)
}
