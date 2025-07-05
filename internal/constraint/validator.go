package constraint

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"haircompany-shop-rest/pkg/response"
)

func ValidateDTO(dto interface{}) []response.ErrorField {
	var errorFields []response.ErrorField
	validate := validator.New()

	err := validate.RegisterValidation("hex_color", validateHexColor)
	if err != nil {
		errorFields = append(errorFields, response.NewErrorField("validation", string(response.ServerError)))
		return errorFields
	}

	err = validate.Struct(dto)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, ve := range validationErrors {
				fieldName := ve.Field()
				tag := ve.Tag()
				errorCode := response.GetErrorCodeByTag(tag)

				errorFields = append(errorFields, response.NewErrorField(fieldName, string(errorCode)))
			}
		}
	}

	return errorFields
}
