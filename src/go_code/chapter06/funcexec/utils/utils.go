package utils

import "fmt"

func PrintPyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		// 在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		// j 表示每层打印多少个*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
