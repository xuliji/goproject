package model

// 定义一个结构体
type student struct {
	Name   string
	Scores float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:   n,
		Scores: s,
	}
}
