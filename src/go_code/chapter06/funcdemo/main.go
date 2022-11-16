package main

import (
	"fmt"
)

func cal(n1 float64, n2 float64, operater byte) float64 {
	var res float64
	switch operater {
	case '-':
		res = n1 - n2
	case '+':
		res = n1 + n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("input error")
	}
	return res
}

func main() {
	var n1, n2 = 20.0, 30.0
	var operater byte = '*'
	res := cal(n1, n2, operater)
	fmt.Println("res =", res)
}
