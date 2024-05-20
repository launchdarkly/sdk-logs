package logs

import (
	"reflect"
	"testing"
)

func TestParseEmptyMessage(t *testing.T) {
	parsed, err := ParseMessage("")
	if err != nil {
		t.Fatal("unexpected error parsing empty message")
	}
	if parsed.Message != "" {
		t.Fatal("parsed message has unexpected content")
	}
	if len(parsed.Parameters) != 0 {
		t.Fatal("parser found parameters that don't exist")
	}
}

func TestParseMessageWithSeveralParameters(t *testing.T) {
	message := "Config option \"${name}\" should be of type ${expectedType}, but received ${actualType}, using the default value (${defaultValue})."
	parsed, err := ParseMessage(message)
	if err != nil {
		t.Fatalf("unexpected error parsing empty message, received: %s", err)
	}
	if parsed.Message != message {
		t.Fatalf("parsed message has unexpected content, expected: %s, actual: %s", message, parsed.Message)
	}
	if len(parsed.Parameters) != 4 {
		t.Fatal("parser did not find expected number of parameters")
	}
	expectedParameters := []string{"name", "expectedType", "actualType", "defaultValue"}
	if !reflect.DeepEqual(parsed.Parameters, expectedParameters) {
		t.Fatalf("did not get expected parameters list, expected: %v, received: %v", expectedParameters, parsed.Parameters)
	}
}

func TestMessageEndsInEscape(t *testing.T) {
	_, err := ParseMessage("\\")
	if err == nil {
		t.Fatal("expected an error and there was not one")
	}
	if err.Error() != "message string ended in escape character" {
		t.Fatalf("did not get the correct error for a message that ends in an escape character, recieved: %s", err.Error())
	}
}

func TestMessageEndsWithOpenParameter(t *testing.T) {
	_, err := ParseMessage("${test")
	if err == nil {
		t.Fatal("expected an error and there was not one")
	}
	expectedError := "message string ended in an open parameter"
	if err.Error() != expectedError {
		t.Fatalf("did not get the correct error for a message that ends in an escape character, recieved: %s", err.Error())
	}

	_, err = ParseMessage("${")
	if err == nil {
		t.Fatal("expected an error and there was not one")
	}
	if err.Error() != expectedError {
		t.Fatalf("did not get the correct error for a message that ends in an escape character, recieved: %s", err.Error())
	}
}

func TestParameterNameContainsInvalidCharacter(t *testing.T) {
	_, err := ParseMessage("${te st}")
	if err == nil {
		t.Fatal("expected an error and there was not one")
	}
	expectedError := "invalid character in message parameter must be only ASCII letters"
	if err.Error() != expectedError {
		t.Fatalf("did not get the correct error for a message that ends in an escape character, recieved: %s", err.Error())
	}
}
