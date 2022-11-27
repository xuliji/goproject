package main

import "fmt"

type AInterface interface {
	Say()
}
type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("我是学生.....")
}

func main() {
	var stu Student // 结构体变量 实现了Say方法
	// 接口自身不能实例化 但是接口可以指向一个实现了该接口的结构体变量
	var a AInterface = stu
	a.Say()
}
