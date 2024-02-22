package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type mathStruct struct {
	number    float64
	mathError error
}

func (e mathStruct) Error() string {
	return fmt.Sprintf("Error: %v", e.mathError)
}

func main() {
	mathStruct1 := mathStruct{
		number:    -12.00,
		mathError: errors.New("number below 0"),
	}
	result, err := sqrtRoot(mathStruct1)
	if err != nil {
		log.Fatalf("error in sqrtRoot: %v", err)
	}
	fmt.Printf("Operation successfull. Result: %v", result)
}

func sqrtRoot(mathStruct1 mathStruct) (float64, error) {
	if mathStruct1.number < 0 {
		return 0, mathStruct1.mathError
	}
	return math.Sqrt(mathStruct1.number), nil
}
