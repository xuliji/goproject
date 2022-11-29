package main

import (
	"fmt"
	"os"
)

// os.Args() 返回一个[]string
func main() {
	fmt.Println("命令行参数有多少个:", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("第%d个参数为%v\n", i+1, arg)
	}
}
