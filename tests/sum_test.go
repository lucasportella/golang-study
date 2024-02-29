package tests

import "testing"

func TestSum(t *testing.T) {
	expected := 6
	result := Sum(2, 3, 1)
	if result != expected {
		t.Error("Expected:", result)
	}
}
