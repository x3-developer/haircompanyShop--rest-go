package auth

import (
	"haircompany-shop-rest/internal/modules/v1/auth/dto"
	"haircompany-shop-rest/internal/modules/v1/client_user"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user"
	"haircompany-shop-rest/internal/services"
)

type Service interface {
	DashboardLogin(loginDto dto.DashboardLoginDTO) (*dto.ResponseDTO, error)
}

type service struct {
	jwtSvc            services.JWTService
	passwordSvc       services.PasswordService
	dashboardUserRepo dashboard_user.Repository
	clientUserRepo    client_user.Repository
}

func NewService(jwtSvc services.JWTService, passwordSvc services.PasswordService, dashboardUserRepo dashboard_user.Repository, clientUserRepo client_user.Repository) Service {
	return &service{
		jwtSvc:            jwtSvc,
		passwordSvc:       passwordSvc,
		dashboardUserRepo: dashboardUserRepo,
		clientUserRepo:    clientUserRepo,
	}
}

func (s *service) DashboardLogin(loginDto dto.DashboardLoginDTO) (*dto.ResponseDTO, error) {
	user, err := s.dashboardUserRepo.GetByEmail(loginDto.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	if err = s.passwordSvc.CompareHashAndPassword(user.Password, loginDto.Password); err != nil {
		return nil, nil
	}

	token, err := s.jwtSvc.GenerateDashboardToken(user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.ResponseDTO{
		Token:            token,
		RefreshToken:     "",
		RefreshExpiresAt: 0,
	}, nil
}
