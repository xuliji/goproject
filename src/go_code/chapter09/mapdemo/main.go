package main

import "fmt"

// map简单介绍
// map的值类型也可以是结构体和map
// map需要make才能使用 make()是用来分配空间的
// map的键值对是乱序的
func main() {
	var map1 map[string]string
	fmt.Println(map1)
	a := make(map[string]string, 10)
	a["no1"] = "love"
	a["no2"] = "you"
	a["no1"] = "I"
	fmt.Println(a)

	cities := map[string]string{
		"no4": "成都",
	}
	fmt.Println(cities)
}
