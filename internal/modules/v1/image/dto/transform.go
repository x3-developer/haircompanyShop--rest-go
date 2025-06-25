package dto

func TransformImageToResponseDTO(fileName string) *ResponseDTO {
	return &ResponseDTO{
		Image: fileName,
	}
}
