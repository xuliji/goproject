package main

import (
	"fmt"
	"strconv"
)

// string转基本数据类型，基本数据类型转string
func main() {
	var num1 = 99
	var num2 float64 = 23.456
	b := true
	//var mychar byte = 'h'
	var str string

	//使用第一种方式进行转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T \nstr=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T \nstr=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T \nstr=%v\n", str, str)

	//第二种 strconv 包中的函数
	var num3 int64 = 99
	var num4 float64 = 23.456
	b1 := true
	str = strconv.FormatInt(num3, 10)
	fmt.Printf("str type %T \nstr=%v\n", str, str)
	//说明：'f'：格式 10：表示小数位保留10位 64：表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T \nstr=%v\n", str, str)

	str = strconv.FormatBool(b1)
	fmt.Printf("str type %T \nstr=%q\n", str, str)
}
