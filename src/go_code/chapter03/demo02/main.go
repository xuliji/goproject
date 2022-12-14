package main

import (
	"fmt"
	"reflect"
)

// 一次性申明多个全局变量
var (
	n8    = 100
	n9    = 11.11
	name3 = "peter"
)

func main() {
	// golang变量使用方式1
	// 指定变量类型，声明后不赋值，使用默认值
	var i int
	fmt.Println(i)
	var j string
	fmt.Println(j)

	// 方式2：根据变量的值自行推导数据类型
	var a = 10.10
	fmt.Println(a)

	// 方式3：省略var，:=左侧的变量不应该是已经申明过的，否则编译不通过
	//等价于var name string   name = "tom"
	name := "tom"
	fmt.Println(name)

	// 方法4：多变量申明
	// 方式1
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	// 方式2
	var n4, name_1, n5 = 10, "tom", 10.11
	fmt.Println(n4, name_1, n5)

	// 方式3
	n6, name_2, n7 := 10, "jack", 13.14
	fmt.Println(n6, name_2, n7)
	fmt.Println(n8, name3, n9)

	// 程序中+号的应用
	// 左右是字符串表示拼接，是数字表示加法
	str1, str2 := "tom ", "love peter"
	str3 := str1 + str2
	fmt.Println(str3)

	// 类型相同才能相加
	var number1, number2 = 10.0, 20.10
	number3 := number2 + number1
	fmt.Println(reflect.TypeOf(number3), reflect.TypeOf(str3))
	fmt.Printf("number3类型 %T\n", number3)
}
