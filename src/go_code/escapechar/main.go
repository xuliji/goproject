package main

/*
常用的转义字符
\t 制表符，用来排版
\n 换行符
\r 一个回车  回车后不换行，相当于替换了回车前的字符串
*/
import "fmt"

func main() {
	fmt.Println("tomjack")
	fmt.Println("tom\tjack")
	fmt.Println("\\\\")
	fmt.Println("我爱是sb\r你们")
	fmt.Println("tom说\"i love you\"")
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}
