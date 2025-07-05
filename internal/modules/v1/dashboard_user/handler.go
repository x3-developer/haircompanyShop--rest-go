package dashboard_user

import (
	"fmt"
	_ "haircompany-shop-rest/docs/response"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user/dto"
	"haircompany-shop-rest/pkg/request"
	"haircompany-shop-rest/pkg/response"
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

// Create creates a new dashboard user
//
//	@Summary		Create a new dashboard user
//	@Description	Create a new dashboard user with the provided details.
//	@Tags			Dashboard User
//	@Security		BearerAuth
//	@Param			Authorization	header		string						true	"Bearer {token}"
//	@Accept			json
//	@Produce		json
//	@Param			dashboardUser	body		dto.CreateDTO						true	"Dashboard User Create DTO"
//	@Success		201				{object}	docsResponse.DashboardUserCreate201	"Dashboard User Created"
//	@Failure		400				{object}	docsResponse.DashboardUserCreate400	"Bad Request or Validation Error"
//	@Failure		500				{object}	docsResponse.Response500			"Server Error"
//	@Router			/api/v1/dashboard-user/create [post]
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

	createdUser, errFields, err := h.svc.Create(createDto)
	if err != nil {
		msg := fmt.Sprintf("failed to create user: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	response.SendSuccess(w, http.StatusCreated, createdUser)
}
