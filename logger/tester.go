package logger

import "testing"

func PrintLogStringResult(functionSignature string, actualValue string, expectedValue string, t *testing.T) {
	if actualValue != expectedValue {
		t.Errorf("%v failed, expected %v, got %v", functionSignature, expectedValue, actualValue)
	}

	t.Logf("%v success", functionSignature)
}

func PrintLogIntResult(functionSignature string, actualValue int, expectedValue int, t *testing.T) {
	if actualValue != expectedValue {
		t.Errorf("%v failed, expected %v, got %v", functionSignature, expectedValue, actualValue)
	}

	t.Logf("%v success", functionSignature)
}
