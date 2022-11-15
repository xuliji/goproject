package main

import (
	"fmt"
	util "go_code/chapter06/fundemo01/utils" // 包可以取别名，但是原先的名字不能用了
)

func main() {
	var n1, n2 = 20.0, 30.0
	var operator byte = '*'

	res := util.Cal(n1, n2, operator)
	fmt.Println("计算结果是：", res)
	util.Print_hello()
}
