package main

import "fmt"

func main() {

	// string底层是byte数组，因此可以进行切片处理
	str := "hello@atguigu"
	// 使用切片获取到atguigu
	slice := str[6:]
	fmt.Println(slice)

	// 修改字符串
	// 转中文不行  中文先转成[]rune  []rune按字符处理，兼容汉字
	arr := []rune(str)
	arr[0] = '爱'
	str = string(arr)
	fmt.Println(str)

}
