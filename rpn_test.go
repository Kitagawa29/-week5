package main

import "testing"

//func TestCalculate01(t *testing.T) {
	//num := Calculate(1,1,"+")
	//if num != 42 {
		//t.Error("Test01 is failed")
	//}
//}

func TestCalculate02(t *testing.T) {
	num := Calculate(2,3,"+")
	if num != 5 {
		t.Error("Test02 is failed")
	}
}

