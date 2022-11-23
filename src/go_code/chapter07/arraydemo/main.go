package main

import "fmt"

func main() {
	// 定义数组
	var hens [6]float64
	// 给元素赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	total_weight := 0.0
	// 遍历数组
	for i := 0; i < len(hens); i++ {
		total_weight += hens[i]
	}
	avg_weight := fmt.Sprintf("%.2f\n", total_weight/float64(len(hens)))
	fmt.Println("avg_weight=", avg_weight)
}
