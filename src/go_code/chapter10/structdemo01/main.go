package main

import "fmt"

// 结构体相当于class  大写开头表示是公开的
// 当申明或者创建一个结构体对象时,空间已经被分配
// 还可以绑定方法
// 结构体是值类型
type Cats struct {
	name   string
	age    int
	color  string
	Hobby  string
	Scores [10]int
}

func main() {
	cat := Cats{
		name:  "小白",
		age:   3,
		color: "花色",
		Hobby: "吃鱼",
	}
	fmt.Println(cat)
}
