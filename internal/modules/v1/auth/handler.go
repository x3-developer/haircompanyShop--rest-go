package auth

import (
	"fmt"
	_ "haircompany-shop-rest/docs/response"
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

//	@Summary		Dashboard user login
//	@Description	Authenticate dashboard user with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		dto.DashboardLoginDTO			true	"Login credentials"
//	@Success		200			{object}	docsResponse.DashboardLogin200	"Login successful"
//	@Failure		400			{object}	docsResponse.DashboardLogin400	"Bad Request or Validation Error"
//	@Failure		401			{object}	docsResponse.Response401		"Unauthorized"
//	@Failure		500			{object}	docsResponse.Response500		"Server Error"
//	@Router			/api/v1/auth/dashboard/login [post]
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

//	@Summary		Dashboard refresh token
//	@Description	Refresh authentication token for dashboard user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		dto.RefreshTokenDTO						true	"Refresh token"
//	@Success		200				{object}	docsResponse.DashboardRefreshToken200	"Token refreshed successfully"
//	@Failure		400				{object}	docsResponse.DashboardRefreshToken400	"Bad Request or Validation Error"
//	@Failure		401				{object}	docsResponse.Response401				"Unauthorized or Invalid Token"
//	@Failure		500				{object}	docsResponse.Response500				"Server Error"
//	@Router			/api/v1/auth/dashboard/refresh-token [post]
func (h *Handler) DashboardRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshTokenDto, err := request.DecodeBody[dto.RefreshTokenDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		response.SendError(w, http.StatusBadRequest, msg, response.BadRequest)
		return
	}

	errFields := constraint.ValidateDTO(refreshTokenDto)
	if errFields != nil {
		msg := "validation errors occurred"
		response.SendValidationError(w, http.StatusBadRequest, msg, response.BadRequest, errFields)
		return
	}

	tokenPair, err := h.svc.DashboardRefreshToken(refreshTokenDto)
	if err != nil {
		msg := fmt.Sprintf("failed to refresh token: %v", err)
		response.SendError(w, http.StatusUnauthorized, msg, response.Unauthorized)
		return
	}

	response.SendSuccess(w, http.StatusOK, tokenPair)
}
