package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// TranslateErrorMessage menerjemahkan pesan error dari validasi dan GORM
func TranslateErrorMessage(err error) map[string]string {

	errorsMap := make(map[string]string)

	// Handle validation errors dari validator v10
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = fmt.Sprintf("%s must be a valid email address", field)
			case "Unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters long", field, fieldError.Param())
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters long", field, fieldError.Param())
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			default:
				errorsMap[field] = "Invalid value for"
			}
		}
	}

	// Handle GORM errors
	if err != nil {
		// Cek jika error mengandung "Duplicate entry" (duplikasi data di database)
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {

			// Jika data yang dicari tidak ditemukan di database
			errorsMap["record"] = "Record not found"
		}
	}

	return errorsMap
}

// IsduplicateEntryError checks if the error is a duplicate entry error
func IsDuplicateEntryError(err error) bool {
	// Cek apakah error mengandung pesan "duplicate entry
	return err != nil && strings.Contains(err.Error(), "duplicate entry")
}
