package dashboard_user

import (
	"haircompany-shop-rest/internal/modules/v1/dashboard_user/dto"
	"haircompany-shop-rest/internal/services"
	"haircompany-shop-rest/pkg/response"
)

type Service interface {
	Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error)
}

type service struct {
	repo        Repository
	passwordSvc services.PasswordService
}

func NewService(r Repository, passwordSvc services.PasswordService) Service {
	return &service{
		repo:        r,
		passwordSvc: passwordSvc,
	}
}

func (s *service) Create(createDto dto.CreateDTO) (*dto.ResponseDTO, []response.ErrorField, error) {
	var validationErrors []response.ErrorField
	existingUser, err := s.repo.GetByEmail(createDto.Email)
	if err != nil {
		return nil, nil, err
	}
	if existingUser != nil {
		if existingUser.Email == createDto.Email {
			validationErrors = append(validationErrors, response.NewErrorField("email", string(response.NotUnique)))
		}

		return nil, validationErrors, nil
	}

	userModel := dto.TransformCreateDTOToModel(createDto)
	passwordHash, err := s.passwordSvc.GenerateHash(createDto.Password)
	if err != nil {
		return nil, nil, err
	}
	userModel.Password = passwordHash

	createdUser, err := s.repo.Create(userModel)
	if err != nil {
		return nil, nil, err
	}

	createdUserResponse := dto.TransformModelToResponseDTO(createdUser)

	return createdUserResponse, nil, nil
}
