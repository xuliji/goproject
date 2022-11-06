package main

import "fmt"

func main() {
	var x, y = 4, 1
	if x > 2 {
		if y > 2 {
			fmt.Println(x + y)
		}
		fmt.Println("atguigu")
	} else {
		fmt.Println("x is =", x)
	}
}
