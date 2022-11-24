package main

import "fmt"

func main() {
	var scores [2][2]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			_, _ = fmt.Scanln(&scores[i][j])
		}
	}
	// 统计平均分
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		fmt.Printf("班级%v的总分为%v, 平均分为%v\n", i, sum, sum/float64(len(scores[i])))
	}
}
