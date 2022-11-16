package main

import (
	"fmt"
	"funcexec/utils"
)

func main() {
	var num_layer int
	fmt.Println("请输入层数：")
	fmt.Scanln(&num_layer)
	utils.PrintPyramid(num_layer)
}
