package advalidator

import (
	"reflect"
)

type ValidationErrors struct {
	Field   string
	Message string
}

func ValidateAd(ad interface{}) []ValidationErrors {
	var errors []ValidationErrors
	adValue := reflect.ValueOf(ad)
	adType := adValue.Type()
	for i := 0; i < adValue.NumField(); i++ {
		fieldValue := adValue.Field(i)
		field := adType.Field(i)

		if field.Name == "Title" {
			value := fieldValue.String()
			if len(value) == 0 {
				errors = append(errors, ValidationErrors{"Title", "Title should not be empty"})
			} else if len(value) > 100 {
				errors = append(errors, ValidationErrors{"Title", "Title should not be longer than 100 characters"})
			}
		}

		if field.Name == "Text" {
			value := fieldValue.String()
			if len(value) == 0 {
				errors = append(errors, ValidationErrors{"Text", "Text should not be empty"})
			} else if len(value) > 500 {
				errors = append(errors, ValidationErrors{"Text", "Text should not be longer than 500 characters"})
			}
		}
	}

	return errors
}
