package main

import (
	"fmt"
	"strconv"
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

func Frac(x int, y int) string {
	return strconv.Itoa(x) + "/" + strconv.Itoa(y)
}

func Rpn(s string) int {
	var stack intStack = make([]int, 0)
	a := strings.Fields(s)
	var num int
	for i := range a {
		if strings.Contains("+-*/", a[i]) {
			y := stack.pop()
			x := stack.pop()
			stack.push(Calculate(x, y, a[i]))
		} else {
			num, _ = strconv.Atoi(a[i])
			stack.push(num)
		}
	}
	return stack.pop()
}

func main() {
	// var stack intStack = make([]int, 0)
	// stack.push(1)
	// stack.push(2)
	// fmt.Println(stack.pop())
	// stack.push(3)
	// fmt.Println(stack.pop())
	// fmt.Println(stack.pop())
	fmt.Println(Rpn("1 2 +"))
}
