package main

import "testing"

func TestCalculate01(t *testing.T) {
	num := Calculate(1,1,"+")
	if num != 42 {
		t.Error("Test01 is failed")
	}
}

