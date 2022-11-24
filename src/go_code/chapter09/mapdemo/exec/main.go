package main

import "fmt"

// 课堂案例
func main() {
	students := make(map[string]map[string]string)
	students["01"] = map[string]string{
		"name": "xuliji",
		"sex":  "男",
	}

	students["02"] = map[string]string{
		"name": "marry",
		"sex":  "女",
	}

	fmt.Println(students)
}
