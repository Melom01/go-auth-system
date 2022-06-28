package logger_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomLogger(t *testing.T) {
	var (
		message             = "logger message"
		expectedMessageType = "string"
	)

	if reflect.TypeOf(message) != reflect.TypeOf(expectedMessageType) {
		t.Errorf("Got %v (%T), wanted %T", message, message, expectedMessageType)
	}
}

func TestCustomFatalLogger(t *testing.T) {
	var (
		message             = "fatal logger message"
		err                 = fmt.Errorf("fatal logger error")
		expectedMessageType = "string"
		expectedErrType     = fmt.Errorf("error")
	)

	if reflect.TypeOf(message) != reflect.TypeOf(expectedMessageType) {
		t.Errorf("Got %v (%T), wanted %T", message, message, expectedMessageType)
	}

	if reflect.TypeOf(err) != reflect.TypeOf(expectedErrType) {
		t.Errorf("Got %v (%T), wanted %T", err, err, expectedErrType)
	}
}
