package main

import "fmt"

func main() {
	// 演示map切片的使用
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = map[string]string{
			"name": "黑山老妖",
			"age":  "10000",
		}
	}

	if monsters[1] == nil {
		monsters[1] = map[string]string{
			"name": "牛魔王",
			"age":  "100000",
		}
	}

	// 动态增加map
	monsters = append(monsters, map[string]string{
		"name": "狐狸精",
		"age":  "200",
	})

	fmt.Println(monsters)
}
