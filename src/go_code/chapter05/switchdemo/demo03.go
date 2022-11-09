package main

import "fmt"

// switch 穿透 fallthrought
func main() {
	var num = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配成功")
	}
}
