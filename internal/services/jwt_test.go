package services

import (
	"testing"
)

func TestJWTService_GenerateDashboardTokenPair(t *testing.T) {
	jwtService := NewJWTService("test-dashboard-secret", "test-client-secret")

	email := "test@example.com"
	role := "admin"

	tokenPair, err := jwtService.GenerateDashboardTokenPair(email, role)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tokenPair.AccessToken == "" {
		t.Error("Expected access token to be generated")
	}

	if tokenPair.RefreshToken == "" {
		t.Error("Expected refresh token to be generated")
	}

	if tokenPair.AccessToken == tokenPair.RefreshToken {
		t.Error("Access token and refresh token should be different")
	}
}

func TestJWTService_GenerateClientTokenPair(t *testing.T) {
	jwtService := NewJWTService("test-dashboard-secret", "test-client-secret")

	phone := "+1234567890"

	tokenPair, err := jwtService.GenerateClientTokenPair(phone)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tokenPair.AccessToken == "" {
		t.Error("Expected access token to be generated")
	}

	if tokenPair.RefreshToken == "" {
		t.Error("Expected refresh token to be generated")
	}
}

func TestJWTService_ValidateDashboardToken(t *testing.T) {
	jwtService := NewJWTService("test-dashboard-secret", "test-client-secret")

	email := "test@example.com"
	role := "admin"

	tokenPair, err := jwtService.GenerateDashboardTokenPair(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	claims, err := jwtService.ValidateDashboardToken(tokenPair.AccessToken)
	if err != nil {
		t.Fatalf("Expected no error validating token, got %v", err)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestJWTService_ValidateClientToken(t *testing.T) {
	jwtService := NewJWTService("test-dashboard-secret", "test-client-secret")

	phone := "+1234567890"

	tokenPair, err := jwtService.GenerateClientTokenPair(phone)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	claims, err := jwtService.ValidateClientToken(tokenPair.AccessToken)
	if err != nil {
		t.Fatalf("Expected no error validating token, got %v", err)
	}

	if claims.Phone != phone {
		t.Errorf("Expected phone %s, got %s", phone, claims.Phone)
	}
}

func TestJWTService_ValidateInvalidToken(t *testing.T) {
	jwtService := NewJWTService("test-dashboard-secret", "test-client-secret")

	_, err := jwtService.ValidateDashboardToken("invalid-token")
	if err == nil {
		t.Error("Expected error for invalid token")
	}

	_, err = jwtService.ValidateDashboardToken("")
	if err == nil {
		t.Error("Expected error for empty token")
	}
}

func TestJWTService_ValidateTokenWithWrongSecret(t *testing.T) {
	jwtService1 := NewJWTService("secret1", "client-secret1")
	jwtService2 := NewJWTService("secret2", "client-secret2")

	email := "test@example.com"
	role := "admin"

	tokenPair, err := jwtService1.GenerateDashboardTokenPair(email, role)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	_, err = jwtService2.ValidateDashboardToken(tokenPair.AccessToken)
	if err == nil {
		t.Error("Expected error when validating token with wrong secret")
	}
}
