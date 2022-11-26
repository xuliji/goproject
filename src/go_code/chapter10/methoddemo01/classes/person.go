package classes

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string
	Age   int
	Skill string
}

func (p Person) Test() {
	fmt.Printf("我是%v, 今年%d\n", p.Name, p.Age)
}

// 传参方法和函数一样
func (p Person) Jisuan(n int) (res int) {
	res = n * (n + 1) / 2
	return res
}

// 打印宠物
func (p Person) Pet(pet *Dog) {
	pet.Name = "1111"
	fmt.Printf("我是%v, 我的宠物是%v\n", p.Name, pet.Name)
}
