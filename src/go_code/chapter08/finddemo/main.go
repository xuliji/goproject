package main

import (
	"goproject/src/go_code/chapter08/finddemo/finds"
	"goproject/src/go_code/chapter08/sortdemo/sorts"
)

func main() {
	name_list := []string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}

	finds.OrderFind(name_list, "11111")

	int_slice := []int{12, 46, 11, 9, 0, 1}
	// 先排序
	sorts.Bubble(int_slice)
	// 再查找
	finds.BinaryFind(int_slice, 100, 0, len(int_slice)-1)

}
