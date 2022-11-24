package main

import "fmt"

// map的增删改查
func main() {
	cities := map[string]string{
		"no1": "北京",
		"no2": "上海",
		"no3": "成都",
		"no4": "合肥",
	}
	fmt.Println(cities)
	// 使用内置函数delete()删除键值对  当键值不存在时不会操作,也不会报错
	delete(cities, "no3")
	fmt.Println(cities)

	// 一次性删除所有的key
	// 方法一:遍历map一个个删除
	// 方法二:使用make()新建一个map,旧的让系统回收

	// map查找
	val, findres := cities["no3"]
	if findres {
		fmt.Println("找到了val=", val)
	} else {
		fmt.Println("没找到")
	}

}
