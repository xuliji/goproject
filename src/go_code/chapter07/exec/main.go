package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intArr := [5]int{}
	// 为了每次生成的随机数不一样，要设置一个随机数种子(seed值)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {

		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)

	// 临时变量
	temp := 0
	for i := 0; i < len(intArr)/2; i++ {
		temp = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println(intArr)

}
