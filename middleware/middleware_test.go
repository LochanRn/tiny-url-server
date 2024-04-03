package middleware

import (
	"testing"
)

func TestGetStackErrorResponse(t *testing.T) {
	// Test case 1: when v is a string
	stackTrace := "some stack trace"
	v := "some error message"

	err := GetStackErrorResponse(stackTrace, v)

	// Check if the error message matches the input string
	if err.Error() != v {
		t.Errorf("Expected error message: %s, but got: %s", v, err.Error())
	}

	// Test case 2: when v3 is neither string nor error
	v2 := 123

	err = GetStackErrorResponse(stackTrace, v2)

	// Check if the error message is "server panic"
	expectedErrMsg := "server panic"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message: %s, but got: %s", expectedErrMsg, err.Error())
	}
}
