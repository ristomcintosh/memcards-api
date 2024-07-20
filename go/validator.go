package main

import (
	"errors"
	"fmt"
	"reflect"
)

type ValidationErrors []string

func ValidateRequestBody(s any) (ValidationErrors, error) {
	var validationErrors ValidationErrors
	data := reflect.ValueOf(s)

	if data.Kind() != reflect.Struct {
		return nil, errors.New("expected a struct")
	}

	for i := 0; i < data.NumField(); i++ {
		field := data.Type().Field(i)
		fieldValue := data.Field(i)

		fieldTag := field.Tag.Get("validate")
		if fieldTag == "required" && (fieldValue.IsZero() || fieldValue.String() == "") {
			validationErrors = append(validationErrors, fmt.Sprintf("Field: '%s' is required", field.Name))
		}
	}

	return validationErrors, nil
}
