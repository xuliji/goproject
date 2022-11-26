package main

import (
	"fmt"
	"goproject/src/go_code/chapter10/methoddemo01/classes"
)

// 方法是和类型绑定的
// 自定义类型都可以有方法,不仅仅是结构体
type A struct {
	Num int
}

// (a A)表示test方法是和结构体A绑定的
func (a A) test() {
	fmt.Println(a.Num)
}

func main() {
	a := A{Num: 1000}
	a.test() // 调用方法
	// 方法绑定后只能通过被绑定的类型调用,不能使用其他的方法调用
	p1 := classes.Person{
		Name: "Peter",
		Age:  18,
	}
	p1.Test() // 跨包时方法名也要大写
	res := p1.Jisuan(1000)
	fmt.Println(res)

	dog1 := classes.Dog{
		Name:  "happy",
		Age:   3,
		Skill: "吃屎",
	}
	p1.Pet(&dog1) // 传指针就会改变结构体的值
	fmt.Println(dog1.Name)

	// 传递结构体的指针, 标准写法
	(&dog1).Hobby()

	// 如果结构体绑定了String()方法, 那么fmt.Println()就会默认调用这个变量的String()进行输出
	stu := classes.Student{
		Name: "xuliji",
		Age:  23,
	}

	fmt.Println(&stu)
}
