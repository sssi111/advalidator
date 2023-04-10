package advalidator

import (
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   string
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var errors []string
	for _, err := range v {
		errors = append(errors, err.Err)
	}
	return strings.Join(errors, ", ")
}

func ValidateAd(ad interface{}) ValidationErrors {
	var errors ValidationErrors
	adValue := reflect.ValueOf(ad)
	adType := adValue.Type()
	for i := 0; i < adValue.NumField(); i++ {
		fieldValue := adValue.Field(i)
		field := adType.Field(i)

		if field.Name == "Title" {
			value := fieldValue.String()
			if len(value) == 0 {
				errors = append(errors, ValidationError{"Title", "Title should not be empty"})
			} else if len(value) > 100 {
				errors = append(errors, ValidationError{"Title", "Title should not be longer than 100 characters"})
			}
		}

		if field.Name == "Text" {
			value := fieldValue.String()
			if len(value) == 0 {
				errors = append(errors, ValidationError{"Text", "Text should not be empty"})
			} else if len(value) > 500 {
				errors = append(errors, ValidationError{"Text", "Text should not be longer than 500 characters"})
			}
		}
	}

	return errors
}
