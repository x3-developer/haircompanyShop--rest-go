package auth

import (
	"fmt"
	"haircompany-shop-rest/internal/constraint"
	"haircompany-shop-rest/internal/modules/v1/auth/dto"
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

func (h *Handler) DashboardLogin(w http.ResponseWriter, r *http.Request) {
	dashboardLoginDto, err := request.DecodeBody[dto.DashboardLoginDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := constraint.ValidateDTO(dashboardLoginDto)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	authData, err := h.svc.DashboardLogin(dashboardLoginDto)
	if err != nil {
		msg := fmt.Sprintf("failed to login: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.ServerError)
		return
	}
	if authData == nil {
		msg := "authentication failed, user not found or invalid credentials"
		response.SendError(w, http.StatusUnauthorized, msg, response.Unauthorized)
		return
	}

	response.SendSuccess(w, http.StatusOK, authData)
}
