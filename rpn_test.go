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

func TestCalculate03(t *testing.T) {
	num := Calculate(4,2,"-")
	if num != 2 {
		t.Error("Test03 is failed")
	}
}

func TestCalculate04(t *testing.T) {
	num := Calculate(4,3,"*")
	if num != 12 {
		t.Error("Test04 is failed")
	}
}

func TestRpn01(t *testing.T) {
	num := rpn("1 1 +")
	if num != 2 {
		t.Error("Test05 is failed")
	}
}