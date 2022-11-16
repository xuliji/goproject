package main

import "fmt"

// defer会产生延迟,
func sum(n1 int, n2 int) int {
	// 当执行到defer时不会执行defer后面的语句,只是压入到独立的栈中(defer栈)
	// 当函数执行完毕后,再从defer栈中执行
	defer fmt.Println("ok n1=", n1)
	defer fmt.Println("ok n2=", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)

	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println(res)
}
