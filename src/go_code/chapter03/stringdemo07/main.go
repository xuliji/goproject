package main

import "fmt"

// string类型使用
func main() {
	//string的基本使用
	var str2 string = "北京长城 110 hello world!"
	fmt.Println(str2)
	// 字符串一旦赋值，字符串就不能修改了：在Go中字符串不可变
	/*
		1. 双引号，会识别转义字符
		2. 反引号，以字符串的原生形式输出，包括换行和特殊字符
	*/
	var str3 = "abc\nabc"
	fmt.Println(str3)

	var str4 = `
	package main
	
	import "fmt"
	
	// string类型使用
	func main() {
		//string的基本使用
		var str2 string = "北京长城 110 hello world!"
		fmt.Println(str2)
		// 字符串一旦赋值，字符串就不能修改了：在Go中字符串不可变
		/*
			1. 双引号，会识别转义字符
			2. 反引号，以字符串的原生形式输出，包括换行和特殊字符
		*/
		var str3 = "abc\nabc"
		fmt.Println(str3)
`
	fmt.Println(str4)

	//字符串拼接方式
	var str5 = str2 + str3
	fmt.Println(str5)
}
