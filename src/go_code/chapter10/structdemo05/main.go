package main

import (
	"encoding/json"
	"fmt"
)

// tag的用处
type Monster struct {
	Name  string `json:"name"` // 结构体标签
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster := Monster{
		Name:  "牛魔王",
		Age:   10000,
		Skill: "芭蕉扇",
	}
	// 将变量序列化为json格式的字符串
	// 返回两个值  第一个值是[]byte
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonMonster))
	}
}
