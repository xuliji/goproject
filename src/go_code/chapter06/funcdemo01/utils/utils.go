package utils

import "fmt"

// 为了跨包使用，需要将C大写，类似其他语言的public(该函数可导出)
func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64
	switch operator {
	case '-':
		res = n1 - n2
	case '+':
		res = n1 + n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入错误")
	}
	return res
}
