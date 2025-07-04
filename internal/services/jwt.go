package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type DashboardCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientCredentials struct {
	Phone string `json:"phone"`
}

type DashboardClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type ClientClaims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type JWTService interface {
	GenerateDashboardTokenPair(email, role string) (*TokenPair, error)
	GenerateClientTokenPair(phone string) (*TokenPair, error)
	ValidateDashboardToken(tokenString string) (*DashboardClaims, error)
	ValidateClientToken(tokenString string) (*ClientClaims, error)
	generateRefreshToken() (string, error)
}

type jwtService struct {
	dashboardSecret []byte
	clientSecret    []byte
}

func NewJWTService(dashboardSecret, clientSecret string) JWTService {
	return &jwtService{
		dashboardSecret: []byte(dashboardSecret),
		clientSecret:    []byte(clientSecret),
	}
}

func (s *jwtService) GenerateDashboardTokenPair(email, role string) (*TokenPair, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &DashboardClaims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.dashboardSecret)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,
	}, nil
}

func (s *jwtService) GenerateClientTokenPair(phone string) (*TokenPair, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &ClientClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.clientSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  tokenString,
		RefreshToken: refreshToken,
	}, nil
}

func (s *jwtService) ValidateDashboardToken(tokenString string) (*DashboardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &DashboardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.dashboardSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*DashboardClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid dashboard token")
}

func (s *jwtService) ValidateClientToken(tokenString string) (*ClientClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		return s.clientSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*ClientClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid client token")
}

func (s *jwtService) generateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
