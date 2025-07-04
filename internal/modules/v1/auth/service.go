package auth

import (
	"fmt"
	"haircompany-shop-rest/internal/modules/v1/auth/dto"
	"haircompany-shop-rest/internal/modules/v1/client_user"
	"haircompany-shop-rest/internal/modules/v1/dashboard_user"
	"haircompany-shop-rest/internal/services"
	"log"
	"time"
)

type Service interface {
	DashboardLogin(loginDto dto.DashboardLoginDTO) (*dto.ResponseDTO, error)
	DashboardRefreshToken(refreshTokenDto dto.RefreshTokenDTO) (*dto.ResponseDTO, error)
}

type service struct {
	redisSvc          services.RedisService
	jwtSvc            services.JWTService
	passwordSvc       services.PasswordService
	dashboardUserRepo dashboard_user.Repository
	clientUserRepo    client_user.Repository
}

func NewService(redisSvc services.RedisService, jwtSvc services.JWTService, passwordSvc services.PasswordService, dashboardUserRepo dashboard_user.Repository, clientUserRepo client_user.Repository) Service {
	return &service{
		redisSvc:          redisSvc,
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

	activeTokenKey := fmt.Sprintf("refresh_token:%s", user.Email)
	activeRefreshToken, err := s.redisSvc.Get(activeTokenKey)
	if err == nil && activeRefreshToken != "" {
		if err := s.redisSvc.Delete(activeRefreshToken); err != nil {
			log.Printf("Failed to delete old refresh token key %s for user %s: %v", activeRefreshToken, user.Email, err)
		}
		if err := s.redisSvc.Delete(activeTokenKey); err != nil {
			log.Printf("Failed to delete user refresh token key %s: %v", activeTokenKey, err)
		}
	}

	tokenPair, err := s.jwtSvc.GenerateDashboardTokenPair(user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	refreshExpiration := 30 * 24 * time.Hour // 30 days

	if err := s.redisSvc.Set(tokenPair.RefreshToken, user.Email, refreshExpiration); err != nil {
		return nil, err
	}

	if err := s.redisSvc.Set(activeTokenKey, tokenPair.RefreshToken, refreshExpiration); err != nil {
		return nil, err
	}

	return &dto.ResponseDTO{
		Token:            tokenPair.AccessToken,
		RefreshToken:     tokenPair.RefreshToken,
		RefreshExpiresAt: time.Now().Add(refreshExpiration).Unix(),
	}, nil
}

func (s *service) DashboardRefreshToken(refreshTokenDto dto.RefreshTokenDTO) (*dto.ResponseDTO, error) {
	userEmail, err := s.redisSvc.Get(refreshTokenDto.RefreshToken) // получаем email по переданному токену
	if err != nil || userEmail == "" {
		return nil, fmt.Errorf("invalid refresh token")
	}

	activeTokenKey := fmt.Sprintf("refresh_token:%s", userEmail)
	activeRefreshToken, err := s.redisSvc.Get(activeTokenKey)
	if err != nil || activeRefreshToken != refreshTokenDto.RefreshToken {
		return nil, fmt.Errorf("refresh token is not active")
	}

	user, err := s.dashboardUserRepo.GetByEmail(userEmail)
	if err != nil || user == nil {
		return nil, fmt.Errorf("user not found")
	}

	tokenPair, err := s.jwtSvc.GenerateDashboardTokenPair(user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	refreshExpiration := 30 * 24 * time.Hour

	if err := s.redisSvc.Delete(refreshTokenDto.RefreshToken); err != nil {
		return nil, err
	}

	if err := s.redisSvc.Set(tokenPair.RefreshToken, user.Email, refreshExpiration); err != nil {
		return nil, err
	}

	if err := s.redisSvc.Set(activeTokenKey, tokenPair.RefreshToken, refreshExpiration); err != nil {
		return nil, err
	}

	return &dto.ResponseDTO{
		Token:            tokenPair.AccessToken,
		RefreshToken:     tokenPair.RefreshToken,
		RefreshExpiresAt: time.Now().Add(refreshExpiration).Unix(),
	}, nil
}
