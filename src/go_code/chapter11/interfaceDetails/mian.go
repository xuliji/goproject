package main

import "fmt"

type AInterface interface {
	Say()
}
type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}

// DInterface 继承了B和C接口
type DInterface interface {
	BInterface
	CInterface
	test03() //要是想实现D接口就要实现A,B和D中所有的方法
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("我是学生.....")
}

func (stu Student) test01() {
	fmt.Println("test01.....")
}
func (stu Student) test02() {
	fmt.Println("test02.....")
}
func (stu Student) test03() {
	fmt.Println("test03.....")
}

func main() {
	var stu Student // 结构体变量 实现了Say方法
	// 接口自身不能实例化 但是接口可以指向一个实现了该接口的结构体变量
	var a AInterface = stu
	a.Say()

	var d DInterface = stu
	d.test01()
	d.test02()
	d.test03()
}
