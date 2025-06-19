package image

import "haircompany-shop-rest/internal/modules/v1/image/dto"

func TransformImageToResponseDTO(fileName string) *dto.ResponseDTO {
	return &dto.ResponseDTO{
		Image: fileName,
	}
}
