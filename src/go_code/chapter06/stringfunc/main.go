package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello 你好"
	// 统计字符串长度， 按字节返回
	length := len(str)
	fmt.Println("length=", length)

	// 遍历带中文的字符串
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	// 字符串转整数
	// 整数转字符串
	n, err := strconv.Atoi("20")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println(n)
	}

	str2 := strconv.Itoa(30)
	fmt.Printf("str2=%v str2=%T\n", str2, str2)

	// 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("btyes=%v\n", bytes)
	// byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	// 10进制转2，8，16进制 strconv.FormatInt()来完成
	str = strconv.FormatInt(123, 16)
	fmt.Println("123对应的二进制是:", str)

	// 查找字串是不是在指定的字符串中
	flag := strings.Contains("seafood", "foo")
	fmt.Printf("%t\n", flag)

	// 查找一个字符串中字串得个数
	count := strings.Count("abababababa", "aba")
	fmt.Println(count)

	// 不区分大小写字母的字符串比较(==是区分字符大小写的)：strings.EqualFold("abc", "ABC")
	fmt.Println(strings.EqualFold("abc", "ABC"))

	// 返回子串在字符串第一次出现的index值，如果没有返回-1
	fmt.Println(strings.Index("123456789", "789"))

	// 返回子串最后一次出现的位置
	fmt.Println(strings.LastIndex("12233344535456", "2"))

	// 将指定的子串替换成另外的子串，n可以指定你希望替换第几个，如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "go语言", -1)
	fmt.Println(str, str2)

	// 按指定的字符对字符串进行分割,返回数组
	fmt.Println(strings.Split("hello,world,ok", ","))

	// 字符串大小写转换
	fmt.Println(strings.ToLower("ABCDEFG"))
	fmt.Println(strings.ToUpper("abcdefg"))

	// 将字符串左右两边空格去除
	fmt.Println(strings.TrimSpace("   我爱你     "))
	// 去掉制定字符
	fmt.Println(strings.Trim("！   我爱你     ！", "！"))
	// 去掉字符串左边指定字符
	fmt.Println(strings.TrimLeft("！   我爱你     ！", "！ "))
	// 去掉右边指定字符
	fmt.Println(strings.TrimRight("！   我爱你     ！", "！ "))

	// 判断字符串是否已指定字符串开头
	fmt.Println(strings.HasPrefix("ftp://192.168.3.1", "ftp"))
	// 判断字符串是不是已指定字符串结尾
	fmt.Println(strings.HasSuffix("abc我爱你", "你"))

}
