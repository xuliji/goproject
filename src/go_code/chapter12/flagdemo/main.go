package main

import (
	"flag"
	"fmt"
)

// flag包来解析命令行参数
func main() {
	// 定义变量用于接收命令行参数值
	var user, pwd, host string
	var port int

	// 接收用户名
	flag.StringVar(&user, "u", "root", "用户名, 默认为root")
	flag.StringVar(&pwd, "pwd", "123456", "密码, 默认为123456")
	flag.StringVar(&host, "h", "192.168.3.1", "主机, 默认为192.168.3.1")
	flag.IntVar(&port, "p", 8080, "端口号,默认为8080")

	fmt.Println(user, pwd, host, port)
}
