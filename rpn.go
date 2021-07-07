package main

import (
	"fmt"
	"strings"
)

type intStack []int

func (stack *intStack) push(i int) {
	*stack = append(*stack, i)
}

func (stack *intStack) pop() int {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return result
}

func Calculate(x int, y int, e string) int {
	var ans int
	if e == "+" {
		ans = x + y
	} else if e == "-" {
		ans = x - y
	} else if e == "*" {
		ans = x * y
	}
	return ans
}

func Rpn(s string) {
	//var stack intStack = make([]int, 0)
	a := strings.Fields(s)
	fmt.Println(a[1])
	return
}

func main() {
	// var stack intStack = make([]int, 0)
	// stack.push(1)
	// stack.push(2)
	// fmt.Println(stack.pop())
	// stack.push(3)
	// fmt.Println(stack.pop())
	// fmt.Println(stack.pop())
	Rpn("1 2 +")
}
