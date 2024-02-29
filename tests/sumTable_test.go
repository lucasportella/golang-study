package tests

import (
	"fmt"
	"testing"
)

type testData struct {
	input  []int
	answer int
}

func TestSumOnTable(t *testing.T) {
	inputsTable := []testData{
		{input: []int{1, 2, 3}, answer: 6},
		{input: []int{5, 5, 5}, answer: 15},
		{input: []int{10, -3, 2}, answer: 9},
	}

	for _, v := range inputsTable {
		result := Sum(v.input...)
		if result != v.answer {
			t.Errorf("Expected: %v, got: %v", v.answer, result)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 1))
	// Output: 2
}
