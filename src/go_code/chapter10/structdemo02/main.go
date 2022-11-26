package main

import (
	"fmt"
	"goproject/src/go_code/chapter10/structdemo02/classes"
)

func main() {
	num := 10
	person1 := classes.Person{
		Ptr: &num,
		Map1: map[string]string{
			"name": "peter",
			"pwd":  "111111",
		},
		Slice: []int{1, 2, 3},
	}
	fmt.Println(person1)

	// 创建对象  方式三
	var p2 *classes.Person = new(classes.Person) //new出来的是指向person对象的指针
	fmt.Printf("p2=%v  type=%T\n", p2, p2)
	// 因为p2是指针所以赋值需要先取值
	(*p2).Ptr = &num
	(*p2).Map1 = map[string]string{
		"name": "1111",
	}
	(*p2).Slice = []int{1, 2, 3, 4}
	fmt.Println(p2)

}
