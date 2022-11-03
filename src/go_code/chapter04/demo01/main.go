package main

import "fmt"

func main() {
	fmt.Printf("10/4 = %v\n", 10/4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10 / 4.0
	fmt.Println(n2)
	num1, num2 := 20, 3.0
	var num3 = float64(num1) / num2
	fmt.Println(num3)

	// 演示%的使用
	// 取模公式  a % b = a - a / b * b
	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)
}
