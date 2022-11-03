package main

import "fmt"

/*
联系获取num的地址,通过指针变量ptr改变num的值
*/
func main() {
	var num int = 100
	fmt.Printf("%v\n", &num)
	var ptr = &num
	*ptr = 300
	fmt.Printf("%v", num)
}
