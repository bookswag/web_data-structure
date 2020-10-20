package calc

import "testing"

func TestMathAdd(t *testing.T) {
	_, result := Add(1, 2, 3)

	expectedValue := 6

	if result != expectedValue {
		t.Errorf("Add(1,2,3) FAILED, expected %v but got value %v", expectedValue, result)
	}

	t.Logf("Add(1,2,3) SUCCESS")
}

func TestMathAddJustOne(t *testing.T) {
	_, result := Add(1)

	// how can I catch the exception in test method?
	if result != 0 {
		t.Errorf("Add(1) FAILED, expected %v but got value %v", 0, result)
	}

	t.Logf("Add(1) SUCCESS")
}
