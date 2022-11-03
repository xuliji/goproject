package main

import "fmt"

// 关系运算符: 返回结果都是bool类型
// 关系表达式 经常用在if结构的条件中或循环结构的条件中
func main() {
	n1, n2 := 9, 10

	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)

}
