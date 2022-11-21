package main

import (
	"fmt"
	"goproject/src/go_code/chapter06/timefunc/utils"
	"time"
)

// 日期和时间相关函数
func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v  type=%T\n", now, now)

	// 获取到其他的日期格式
	fmt.Printf("年=%v\n月=%v\n日=%v\n", now.Year(), int(now.Month()), now.Day())

	// 日期和时间格式化
	// 方式一：使用Printf()或则Sprintf()
	// 方式二：time.Format()  后面的日期必须是这个时间
	str := fmt.Sprintln(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Printf("%v", str)
	utils.Sleep(3 * time.Second)
	fmt.Println("11111")

	// unix时间戳和unixnano时间戳(用来获取随机数字)
	fmt.Printf("unix时间戳：%v unixnano时间戳：%v\n", now.Unix(), now.UnixNano())
	f := utils.Test("hello")
	strat := time.Now().Unix()
	f(100000)
	coast := time.Now().Unix() - strat
	fmt.Printf("执行test()耗费了%v秒\n", coast)
}
