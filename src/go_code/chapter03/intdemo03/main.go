package main

import (
	"fmt"
	"unsafe"
)

// 整数类型的使用
func main() {
	//uint8不能表示负数，如果没指定数据类型则默认是int32
	var i uint8 = 255
	fmt.Println(i)

	//int表示有符号整数，占多少字节和你的系统相关
	//32位 4字节
	//64位 8字节
	/*
		rune 有符号 与int32一样,表示一个unicode码
		byte 无符号 与uint8等价,储存字符时使用byte
	*/
	var a int = 8900
	var b uint = 10
	var c rune = 666
	var d byte = 1
	//fmt.Printf()用于做格式化输出
	//unsafe.Sizeof() 是unsafe包的函数可以返回变量的占用的字节数
	fmt.Printf("a占%d字节\n", unsafe.Sizeof(a))
	fmt.Printf("b占%v字节\n", unsafe.Sizeof(b))
	fmt.Printf("c占%v字节\n", unsafe.Sizeof(c))
	fmt.Printf("d占%v字节\n", unsafe.Sizeof(d))

	/*
		golang程序中整型变量使用时,遵守保小不保大的原则,
		即:在保证程序正确运行的情况下,尽量使用占用空间小的数据类型
	*/
	var age byte = 100
	fmt.Println(age)

}
