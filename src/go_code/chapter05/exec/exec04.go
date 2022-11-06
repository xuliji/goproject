package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("请输入a,b,c:")
	fmt.Scanf("%f,%f,%f", &a, &b, &c)
	m := b*b - 4*a*c
	if m > 0 {
		var x1, x2 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a), (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
		fmt.Printf("x1=%v, x2=%v", x1, x2)
	} else if m == 0 {
		var x1 = (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Println("方程没有实数解")
	}
}
