package main

import "fmt"

func main() {
	/*
		8s进入决赛，否则淘汰。进入决赛后还要分成男女组
	*/
	var score float64
	var sex string

	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	if score < 8.0 {
		fmt.Println("请输入性别：")
		fmt.Scanln(&sex)
		if sex == "男" {
			fmt.Println("进入决赛男子组")
		} else {
			fmt.Println("进入决赛女子组")
		}
	} else {
		fmt.Println("很遗憾，没进入决赛。")
	}
}
