package validator

import (
	"fmt"
	"strings"

	"github.com/FauzanAr/go-init/pkg/wrapper"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func formatValidationErrors(errs validator.ValidationErrors) error {
	var messages []string

	for _, err := range errs {
		var message string
		param := err.Param()
		fieldName := err.Field()
		tag := err.Tag()

		message = fmt.Sprintf("[%s] must %s", fieldName, tag)
		if param != "" {
			message = message + fmt.Sprintf(" %s", param)
		}

		messages = append(messages, message)
	}

	return wrapper.ValidationError(strings.Join(messages, ", "))
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return formatValidationErrors(err.(validator.ValidationErrors))
	}

	return nil
}
