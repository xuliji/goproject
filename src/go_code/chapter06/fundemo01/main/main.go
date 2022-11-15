package main

import (
	"../utils"
	"fmt"
)

func main() {
	var n1, n2 = 20.0, 30.0
	var operator = '*'

	res := utils.Cal(n1, n2, operator)
	fmt.Println("计算结果是：", res)
}
