package main

// golang错误处理机制  defer, panic, recover结合使用
import (
	"errors"
	"fmt"
)

func test() {
	// 使用defer + recover 来处理异常
	defer func() {
		err := recover() // recover()可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
func readConf(name string) (err error) {
	if name == "config.ini" {
		fmt.Println("读取正确")
		return nil
	} else {
		// 返回自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	if err := readConf("config2.ini"); err != nil {
		// 读取文件发生错误，使用panic打印错误信息并终止程序
		panic(err)
	}
	fmt.Println("继续执行")
}

func main() {
	test()

	fmt.Println("main()函数后面的代码")

	// 自定义错误
	// errors.New
	test02()

}
