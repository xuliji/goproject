package main

import "fmt"

func main() {
	var grade int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&grade)
	if grade == 100 {
		fmt.Println("奖励BMW一台")
	} else if grade > 80 && grade <= 99 {
		fmt.Println("奖励一台iphone7")
	} else if grade >= 60 && grade <= 80 {
		fmt.Println("奖励一台ipad")
	} else {
		fmt.Println("什么奖励都没有")
	}
}
