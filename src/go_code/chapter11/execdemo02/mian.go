package main

import (
	"fmt"
	"goproject/src/go_code/chapter11/execdemo02/classes"
	"sort"
)

// 实现结构体切片排序
func main() {
	heros := classes.HeroSlice{classes.Heros{
		Name: "超人",
		Age:  20,
	}, classes.Heros{
		Name: "钢铁侠",
		Age:  10,
	}, classes.Heros{
		Name: "黑寡妇",
		Age:  30,
	}, classes.Heros{
		Name: "绿巨人",
		Age:  25,
	}}
	sort.Sort(heros)
	fmt.Println(heros)
}
