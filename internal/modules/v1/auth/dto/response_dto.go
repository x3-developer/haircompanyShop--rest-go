package dto

type ResponseDTO struct {
	Token            string `json:"token"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresAt int64  `json:"refreshExpiresAt"`
}
