package utils

import (
	"strings"
)

func MakeSuffix(suffix string) func(string) string {
	// 返回的函数和suffix构成了闭包
	return func(name string) string {
		// 如果没有指定的后缀，则加上，否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}
