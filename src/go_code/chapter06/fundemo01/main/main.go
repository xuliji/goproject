package main

import (
	"fmt"
	"src/go_code/chapter06/fundemo01/utils"
)

func main() {
	var n1, n2 = 20.0, 30.0
	var operator byte = '*'

	res := utils.Cal(n1, n2, operator)
	fmt.Println("计算结果是：", res)
}
