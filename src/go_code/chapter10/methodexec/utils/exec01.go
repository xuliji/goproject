package utils

import "fmt"

type MethodUtils struct {
}

func (m MethodUtils) Print_area() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
