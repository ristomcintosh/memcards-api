package main

import (
	"reflect"
	"testing"
)

func Test_Request_Validator(t *testing.T) {
	tests := []struct {
		data                 any
		expectedErrMsg       ValidationErrors
		expectedErrMsgLength int
	}{
		{struct {
			Front string `validate:"required"`
			Back  string `validate:"required"`
		}{Front: "Test"}, []string{"Field: 'Back' is required"}, 1},
		{struct {
			Front string `validate:"required"`
			Back  string `validate:"required"`
		}{Front: "Test", Back: "Test"}, []string{}, 0},
		{struct {
			Front string `validate:"required"`
			Back  string `validate:"required"`
		}{Front: "", Back: ""}, []string{"Field: 'Front' is required", "Field: 'Back' is required"}, 2},
		{struct{ OptionalField string }{OptionalField: ""}, []string{}, 0},
	}

	for _, tt := range tests {
		result, err := ValidateRequestBody(tt.data)

		if err != nil {
			t.Fail()
		}

		if len(result) != tt.expectedErrMsgLength {
			t.Errorf("Expected to have %d error(s)", tt.expectedErrMsgLength)
		}

		if len(result) != len(tt.expectedErrMsg) && (!reflect.DeepEqual(result, tt.expectedErrMsg)) {
			t.Errorf("Expected error messages to be: %s, got %s", tt.expectedErrMsg, result)
		}
	}

	t.Run("value isn't a struct", func(t *testing.T) {
		structImposter := "I'm a struct"

		_, err := ValidateRequestBody(structImposter)

		if err == nil {
			t.Error("expected an err")
		}

	})
}
