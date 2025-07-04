package services

import (
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

type JWTService interface {
	GenerateDashboardToken(email, role string) (string, error)
	GenerateClientToken(phone string) (string, error)
	ValidateDashboardToken(tokenString string) (*DashboardClaims, error)
	ValidateClientToken(tokenString string) (*ClientClaims, error)
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

func (s *jwtService) GenerateDashboardToken(email, role string) (string, error) {
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

	return tokenString, err
}

func (s *jwtService) GenerateClientToken(phone string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &ClientClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.clientSecret)

	return tokenString, err
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
