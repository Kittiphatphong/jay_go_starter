package validation

import "github.com/go-playground/validator/v10"

type ErrorResp struct {
	Error string
}

func Validate(request interface{}) []*ErrorResp {
	var errors []*ErrorResp
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResp
			param := ""
			if err.Param() != "" {
				param = " " + err.Param()
			}
			element.Error = err.StructField() + " " + err.Tag() + param
			errors = append(errors, &element)
		}
	}
	return errors
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(myStruct interface{}) string {
	var errorX []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(myStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field() + " " + err.Tag() + " " + err.Param()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errorX = append(errorX, &element)
		}

	}
	if errorX != nil {
		return errorX[0].FailedField
	}
	return ""
}
