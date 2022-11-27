package main

// 继承语法
import (
	"errors"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score map[string]float64
}

func (s *Student) ShowInfos() {
	fmt.Printf("学生名字%v 年龄%d 成绩%v\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(className string, score float64) {
	// 初始化
	if s.Score == nil {
		s.Score = map[string]float64{}
	}
	if score > 0 && score <= 150 {
		s.Score[className] = score
	} else {
		err := errors.New("成绩范围不正确")
		panic(err)
	}
}

// Pupil 小学生 继承 Student类
type Pupil struct {
	Student //匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试.....")
}

// Graduate 大学生 继承Student类
type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试....")
}

func main() {

	// 匿名结构体使用方式
	p := Pupil{Student{
		Name: "小明",
		Age:  10,
	}}
	p.testing()
	p.SetScore("语文", 100.0)
	p.ShowInfos()

	g := Graduate{Student{
		Name: "小许",
		Age:  23,
	}}
	g.testing()
	g.SetScore("语文", 150.0)
	g.ShowInfos()
}
