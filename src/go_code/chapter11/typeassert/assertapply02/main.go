package main

import "fmt"

// 判断变量类型
func TypeJudge(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%d个参数是bool型, 值为%v\n", i, v)
		case string:
			fmt.Printf("第%d个参数是string型, 值为%v\n", i, v)
		case int, int32, int64:
			fmt.Printf("第%d个参数是int型, 值为%v\n", i, v)
		default:
			fmt.Printf("第%d个参数类型未知, 值为%v\n", i, v)
		}
	}
}

func main() {
	n1, n2, n3, n4 := 10, 110.0, "你好", true
	TypeJudge(n1, n2, n3, n4)

}
