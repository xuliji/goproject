package classes

import "fmt"

type OldMonkey struct {
	Name string
}

func (m *OldMonkey) Climbing() {
	fmt.Println(m.Name, "生来会爬树")
}

type YoungMonkey struct {
	OldMonkey // 继承
}

// 实现方法
func (m *YoungMonkey) Flying() {
	fmt.Println(m.Name, "学会了飞翔")
}

// 声明接口
type BirdAble interface {
	Flying()
}

func ShowSkills(able BirdAble) {
	able.Flying()
}
