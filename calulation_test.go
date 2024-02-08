package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(10, 5)
	if result != 15 {
		t.Errorf("Addition Failed, Expected %d and got %d,", 15, result)
	} else {
		fmt.Println("passing")
	}
}

func TestSub(t *testing.T) {
	result := sub(10, 5)
	if result != 5 {
		t.Errorf("Subtraction Failed. Expected %d and got %d", 5, result)
	}
}

func TestMultiply(t *testing.T) {
	result := multiple(10, 5)
	if result != 50 {
		t.Errorf("Multiplication Failed. Expected %d and got %d", 50, result)
	}
}

func TestDivision(t *testing.T) {
	result := division(10, 5)
	if result != 2 {
		t.Errorf("Division Failed. Expected %d and got %d", 2, result)
	}
}
func TestDivideByZero(t *testing.T) {
	result := division(0, 0)
	if result != -1 {
		t.Errorf("Division failed, as Divide by 0 is not expected")
	}
}
