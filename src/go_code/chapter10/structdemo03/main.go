package main

import "fmt"

type Point struct {
	x int
	y int
}
type Rect struct {
	leftup, rightdown Point
}

type Rect2 struct {
	leftup, rightdown *Point
}

func main() {
	r1 := Rect{
		leftup:    Point{1, 2},
		rightdown: Point{4, -3},
	}
	// r1的4个int值在内存中是顺序分布的
	fmt.Println(&r1.leftup.x, &r1.leftup.y, &r1.rightdown.x, &r1.rightdown.y)

	// r2的两个指针的地址是连续的,但是它们指向的地址不一定是连续的
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	r2 := Rect2{
		leftup:    &p1,
		rightdown: &p2,
	}
	fmt.Println(&r2.leftup, &r2.rightdown)
	fmt.Printf("r2.leftup指向的地址:%p r2.rightdown指向的地址%p\n", r2.leftup, r2.rightdown)
}
