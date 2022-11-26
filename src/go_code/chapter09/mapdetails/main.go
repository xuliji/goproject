package main

import "fmt"

// 结构体定义
type Stu struct {
	name    string
	age     int
	address string
}

func modify(map1 map[int]int) {
	map1[10] = 900
}

func main() {
	// map是引用类型
	// 修改后,会直接影响原来的map
	map1 := map[int]int{
		10: 200,
	}
	fmt.Println(map1)
	modify(map1)
	fmt.Println(map1)

	// map的value经常使用结构体
	students := map[int]Stu{}
	students[0] = Stu{name: "许立基", age: 22, address: "上海大学"}
	students[1] = Stu{name: "饶剑锋", age: 22, address: "上海大学"}
	fmt.Println(students)

	for i, stu := range students {
		fmt.Println("学生的编号是:", i)
		fmt.Println("学生的姓名是:", stu.name)
		fmt.Println("学生的年龄是:", stu.age)
		fmt.Println("学生的地址是:", stu.address)
	}
}
