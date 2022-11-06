package main

import "fmt"

func main() {
	var i, j int
	fmt.Println("input i and j:")
	fmt.Scanf("%d %d", &i, &j)
	if i+j >= 50 {
		fmt.Println("hello world!")
	}
}
