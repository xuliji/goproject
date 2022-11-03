package main

import "fmt"

// 从控制台获取信息
func main() {
	// 方式一 fmt.Scanln
	var name string
	var age byte
	var sal float64
	var ispass bool
	//fmt.Println("姓名")
	//fmt.Scanln(&name)
	//
	//fmt.Println("age")
	//fmt.Scanln(&age)
	//
	//fmt.Println("sal")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("ispass")
	//fmt.Scanln(&ispass)
	//
	//fmt.Printf("name=%v \nage=%v \nsal=%v \nispass=%v", name, age, sal, ispass)

	// 方式二 fmt.Scanf 可以按指定格式输入
	fmt.Println("请输入你的姓名, 年龄, 薪水, 是否通过考试(中间用空格隔开)")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &ispass)
	fmt.Printf("name=%v \nage=%v \nsal=%v \nispass=%v", name, age, sal, ispass)

}
