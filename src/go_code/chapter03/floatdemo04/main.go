package main

import (
	"fmt"
	"unsafe"
)

// 小数类型的使用
func main() {
	var price float32 = 89.1234
	fmt.Println(price)

	//浮点数可能有精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	/*
		golang 浮点类型有固定的范围和长度,不受具体的OS操作系统影响
		golang 默认浮点型声明为float64类型
	*/
	num5 := 12.1111
	fmt.Printf("num5的数据类型为:%T, 所占字节数为:%d\n", num5, unsafe.Sizeof(num5))
	//科学计数法
	num6 := 5.1234e2
	num7 := 5.1234e2
	fmt.Println(num6, num7)
}
