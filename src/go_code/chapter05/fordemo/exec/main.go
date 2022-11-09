package main

import "fmt"

func main() {
	var max = 100
	var sum, count int
	for i := 1; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%d sum=%d", count, sum)

	var n int
	fmt.Println("请输入n：")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("%v + %v = %v\n", i, n-i, n)
	}
}
