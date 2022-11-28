package main

import (
	"fmt"
	"goproject/src/go_code/chapter11/encapsulate/model"
)

// 实现封装
func main() {
	p1 := model.NewPerson("xuliji")

	p1.SetAge(23)
	p1.SetSalary(300000.0)
	fmt.Printf("年龄是%v, 工资是%.2f\n", p1.GetAge(), p1.GetSalary())

}
