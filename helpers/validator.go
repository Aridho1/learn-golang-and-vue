package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorsMap := make(map[string]string)

	// fmt.Printf("\nerr: %+v\n", err)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		// fmt.Printf("\nvalidationErrors: %+v\n", validationErrors)
		
		for _, fieldError := range validationErrors {
			field := fieldError.Field()

			switch fieldError.Tag() {
				case "required":
					errorsMap[field] = fmt.Sprintf("%s is required", field)
				case "email":
					errorsMap[field] = "Invalid email format"
				case "unique":
					errorsMap[field] = fmt.Sprintf("%s already exists", field)
				case "min":
					errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
				case "max":
					errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
				case "numeric":
					errorsMap[field] = fmt.Sprintf("%s must be number", field)
				default:
					errorsMap[field] = "Invalid value"
			}
		}
	}

	if err != nil {
		_err := err.Error()
		lowerCaseErr := strings.ToLower(_err)

		// fmt.Printf("\n_err: %+v\n", _err)
		
		if strings.Contains(lowerCaseErr, "duplicate entry") {
			if strings.Contains(_err, "username") {
				errorsMap["username"] = "Username already exists"
			} 
			
			if strings.Contains(_err, "email") {
				errorsMap["email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorsMap["Error"] = "Record not found"
		}
	}

	// fmt.Printf("\nerrorsMap: %+v\n", errorsMap)

	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate Entry")
}