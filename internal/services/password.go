package services

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type passwordService struct{}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (p *passwordService) GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p *passwordService) CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
