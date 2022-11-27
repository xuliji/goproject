package main

import (
	"fmt"
	"goproject/src/go_code/chapter10/factorymode/model"
)

// 结构体名字大写可以跨包使用, 小写不行
// 小写需要工厂模式
func main() {
	stu1 := model.NewStudent("xuliji", 30.33)
	fmt.Println((*stu1))
}
